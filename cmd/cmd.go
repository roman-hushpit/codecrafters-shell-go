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
	env, b := os.LookupEnv(`PATH`)
	if !b {
		fmt.Printf("%s: not found\n", functionName)
		return nil
	}

	for _, path := range strings.Split(env, ":") {
		fullPath := path + "/" + functionName
		if fileInfo, err := os.Stat(fullPath); err == nil {
			if fileInfo.Mode().Perm()&0111 != 0 {
				fmt.Printf("%s is %s\n", functionName, fullPath)
				return nil
			}
		}
	}

	fmt.Printf("%s: not found\n", functionName)
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
