package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	Name       string
	Args       []string
	StdoutFile string
}

type ExecutionContext struct {
	Stdout io.Writer
	Stderr io.Writer
}

type Executable interface {
	Execute(c *Command, ctx *ExecutionContext) error
}

var BuiltinsMap map[string]Executable

func init() {
	BuiltinsMap = map[string]Executable{
		"exit": &ExitCommand{},
		"echo": &EchoCommand{},
		"type": &TypeCommand{},
		"pwd":  &PwdCommand{},
		"cd":   &CdCommand{},
	}
}

type ExitCommand struct {
}

func (e *ExitCommand) Execute(_ *Command, ctx *ExecutionContext) error {
	os.Exit(0)
	return nil
}

type CdCommand struct {
}

func (cd CdCommand) Execute(c *Command, ctx *ExecutionContext) error {
	dir := c.Args[0]
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
		_, err := fmt.Fprintln(ctx.Stderr, fmt.Sprintf("%s: No such file or directory", dir))
		if err != nil {
			return err
		}
	}
	return nil
}

type PwdCommand struct {
}

func (pwd PwdCommand) Execute(_ *Command, ctx *ExecutionContext) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(ctx.Stdout, fmt.Sprintf("%s", dir))
	if err != nil {
		return err
	}
	return nil
}

type TypeCommand struct {
}

func (t TypeCommand) Execute(c *Command, ctx *ExecutionContext) error {
	functionName := c.Args[0]
	if _, ok := BuiltinsMap[functionName]; ok {
		_, err := fmt.Fprintln(ctx.Stdout, fmt.Sprintf("%s is a shell builtin", functionName))
		return err
	}
	executablePath, err := FindExecutable(functionName)
	if err != nil {
		_, err = fmt.Fprintln(ctx.Stdout, fmt.Sprintf("%s", err.Error()))
		if err != nil {
			return err
		}
		return nil
	}
	_, err = fmt.Fprintln(ctx.Stdout, fmt.Sprintf("%s is %s", functionName, executablePath))
	return err
}

type EchoCommand struct {
}

func (echo EchoCommand) Execute(c *Command, ctx *ExecutionContext) error {
	_, err := fmt.Fprintln(ctx.Stdout, strings.Join(c.Args, " "))
	return err
}

type ExternalCommandExecutor struct {
}

func (ext ExternalCommandExecutor) Execute(c *Command, ctx *ExecutionContext) error {
	_, err := FindExecutable(c.Name)
	if err != nil {
		return err
	}
	command := exec.Command(c.Name, c.Args...)
	command.Stderr = ctx.Stderr
	command.Stdout = ctx.Stdout
	err = command.Run()
	if err != nil {
		return err
	}
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
