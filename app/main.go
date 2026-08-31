package main

import (
	"io"
	"strings"

	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/autocomplete"
	"github.com/codecrafters-io/shell-starter-go/executor"
	"github.com/codecrafters-io/shell-starter-go/parser"
	"github.com/codecrafters-io/shell-starter-go/processor"
)

func main() {

	var base = readline.NewPrefixCompleter(
		readline.PcItem("echo"),
		readline.PcItem("exit"),
	)

	l, err := readline.NewEx(&readline.Config{
		Prompt:          "$ ",
		AutoComplete:    &autocomplete.ExtendedCompleter{Inner: base},
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	l.CaptureExitSignal()

	for {
		line, lineError := l.Readline()
		if lineError == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if lineError == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		newParser := parser.NewParser()
		newParser.Parse(line)
		arguments := newParser.Args()
		command := processor.ProcessCommandArgs(arguments)
		executor.ExecuteCommand(command)

	}
}
