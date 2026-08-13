package cmd

import "os"

type Builtins map[string]func() error

var BuiltinsMap = map[string]func() error{
	"exit": exitFunc,
}

func exitFunc() error {
	os.Exit(0)
	return nil
}
