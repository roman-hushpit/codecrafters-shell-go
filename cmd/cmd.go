package cmd

import (
	"fmt"
	"os"
	"strings"
)

type Builtins map[string]func() error

var BuiltinsMap map[string]func(args ...string) error

func init() {
	BuiltinsMap = map[string]func(args ...string) error{
		"exit": exitFunc,
		"echo": echoFunc,
		"type": typeFunc,
	}
}

func typeFunc(args ...string) error {
	functionName := args[0]
	if _, ok := BuiltinsMap[functionName]; ok {
		fmt.Printf("%s is a shell builtin\n", functionName)
		return nil
	}
	executablePath, err := FindExecutable(functionName)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return nil
	}
	fmt.Printf("%s is %s\n", functionName, executablePath)
	return nil
}

func echoFunc(args ...string) error {
	fmt.Println(args[0])
	return nil
}

func exitFunc(args ...string) error {
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
