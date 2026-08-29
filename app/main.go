package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/cmd"
	"github.com/codecrafters-io/shell-starter-go/parser"
)

func main() {
	for {
		fmt.Print("$ ")
		input := bufio.NewScanner(bufio.NewReader(os.Stdin))
		if input.Scan() {
			command := input.Text()
			newParser := parser.NewParser()
			newParser.Parse(command)
			arguments := newParser.Args()
			commandName := arguments[0]
			fnc, ok := cmd.BuiltinsMap[commandName]
			if !ok {
				_, err := cmd.FindExecutable(commandName)
				if err != nil {
					fmt.Printf("%s: command not found\n", command)
					continue
				}

				command := exec.Command(commandName, arguments[1:]...)
				command.Stderr = os.Stderr
				command.Stdout = os.Stdout
				err = command.Run()
				if err != nil {
					continue
				}
			} else {
				err := fnc(arguments[1:]...)
				if err != nil {
				}
			}
		}
	}
}
