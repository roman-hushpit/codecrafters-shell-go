package cmd

import (
	"fmt"
	"os"
	"strings"
)

type Builtins map[string]func() error

var BuiltinsMap map[string]func(commandParameters ...string) error

func init() {
	BuiltinsMap = map[string]func(commandParameters ...string) error{
		"exit": exitFunc,
		"echo": echoFunc,
		"type": typeFunc,
		"pwd":  pwdFunc,
		"cd":   cdFunc,
	}
}

func cdFunc(commandParameters ...string) error {
	dir := commandParameters[0]
	if dir == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		_ = os.Chdir(homeDir)
		return nil
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("%s: No such file or directory\n", dir)
		return nil
	}
	return nil
}

func pwdFunc(_ ...string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", dir)
	return nil
}

func typeFunc(commandParameters ...string) error {
	functionName := commandParameters[0]
	if _, ok := BuiltinsMap[functionName]; ok {
		fmt.Printf("%s is a shell builtin\n", functionName)
		return nil
	}
	executablePath, err := FindExecutable(functionName)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil
	}
	fmt.Printf("%s is %s\n", functionName, executablePath)
	return nil
}

func echoFunc(commandParameters ...string) error {
	fmt.Printf("%s\n", strings.Join(commandParameters, " "))
	return nil
}

func exitFunc(_ ...string) error {
	os.Exit(0)
	return nil
}

func FindExecutable(commandName string) (string, error) {

	env, b := os.LookupEnv(`PATH`)
	if !b {
		return "", fmt.Errorf("%s: not found", commandName)
	}

	for _, path := range strings.Split(env, ":") {
		fullPath := path + "/" + commandName
		if fileInfo, err := os.Stat(fullPath); err == nil {
			if fileInfo.Mode().Perm()&0111 != 0 {
				return fullPath, nil
			}
		}
	}

	return "", fmt.Errorf("%s: not found", commandName)
}
