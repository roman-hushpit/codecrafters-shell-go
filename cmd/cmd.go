package cmd

import (
	"fmt"
	"os"
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
		fmt.Printf("%s is a shall builtin\n", functionName)
		return nil
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
