package autocomplete

import (
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/chzyer/readline"
)

type ExtendedCompleter struct {
	Inner      readline.AutoCompleter
	Instance   *readline.Instance
	tabCount   int
	lastPrefix string
}

func (b *ExtendedCompleter) Do(line []rune, pos int) (newLine [][]rune, length int) {
	if string(line[0:pos]) != b.lastPrefix {
		b.lastPrefix = string(line[0:pos])
		b.tabCount = 1
	} else {
		b.tabCount++
	}
	newLine, length = b.Inner.Do(line, pos)

	if len(newLine) == 1 {
		return newLine, length
	} else if len(newLine) > 1 && b.tabCount > 1 {
		// print candidates ourselves, then hand back the line unchanged
		// so the library redraws prompt + line as if inserting a single match
		names := make([]string, len(newLine))
		for i, r := range newLine {
			names[i] = string(line[0:pos]) + string(r)
		}
		os.Stdout.WriteString("\r\n" + strings.Join(names, "  ") + "\r\n")
		return [][]rune{line[pos:]}, pos
	} else if len(newLine) > 1 && b.tabCount == 1 {
		os.Stdout.Write([]byte{7})
		return [][]rune{}, 0
	}

	if len(newLine) == 0 && pos > 0 {
		env, found := os.LookupEnv(`PATH`)
		if !found {
			os.Stdout.Write([]byte{7})
			return newLine, length
		}
		wordStart := 0
		prefix := line[wordStart:pos]
		possibleExecutables := [][]rune{}
		for _, path := range strings.Split(env, ":") {
			if !dirExists(path) {
				continue
			}
			executables, _ := listExecutables(path)
			for _, executable := range executables {
				if after, ok := strings.CutPrefix(executable, string(prefix)); ok {
					possibleExecutables = append(possibleExecutables, []rune(after))
				}
			}
		}
		slices.SortFunc(possibleExecutables, func(a, b []rune) int {
			return slices.Compare(a, b)
		})

		if len(possibleExecutables) == 1 {
			possibleExecutables[0] = append(possibleExecutables[0], ' ')
			return possibleExecutables, pos - wordStart
		} else if len(possibleExecutables) > 1 && b.tabCount > 1 {
			names := make([]string, len(possibleExecutables))
			for i, r := range possibleExecutables {
				names[i] = string(prefix) + string(r)
			}
			os.Stdout.WriteString("\r\n" + strings.Join(names, "  ") + "\r\n")
			return [][]rune{line[pos:]}, pos
		} else {
			os.Stdout.Write([]byte{7})
			return [][]rune{}, 0
		}
	}

	if len(newLine) == 0 {
		os.Stdout.Write([]byte{7})
	}
	return newLine, length
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return info.IsDir()
}

func listExecutables(dirPath string) ([]string, error) {
	var executables []string
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}

		if isExecutable(info) {
			executables = append(executables, file.Name())
		}
	}

	return executables, nil
}

func isExecutable(info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}
	if runtime.GOOS == "windows" {
		ext := filepath.Ext(info.Name())
		return ext == ".exe" || ext == ".bat" || ext == ".cmd"
	}
	return info.Mode()&0111 != 0
}
