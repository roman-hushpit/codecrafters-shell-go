package cmd

import (
	"fmt"
	"os"
)

type Builtins map[string]func() error

var BuiltinsMap = map[string]func(args ...string) error{
	"exit": exitFunc,
	"echo": echoFunc,
}

func echoFunc(args ...string) error {
	fmt.Println(args[0])
	return nil
}

func exitFunc(args ...string) error {
	os.Exit(0)
	return nil
}
