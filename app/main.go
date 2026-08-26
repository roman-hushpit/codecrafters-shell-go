package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd"
)

func main() {
	for {
		fmt.Print("$ ")
		input := bufio.NewScanner(bufio.NewReader(os.Stdin))
		if input.Scan() {
			command := input.Text()
			arguments := strings.SplitN(command, " ", 2)
			commandName := arguments[0]
			var commandArgumentsString string
			if len(arguments) > 1 {
				commandArgumentsString = arguments[1]
			}
			fnc, ok := cmd.BuiltinsMap[commandName]
			if !ok {
				_, err := cmd.FindExecutable(commandName)
				if err != nil {
					fmt.Printf("%s: command not found\n", command)
					continue
				}
				command := exec.Command(commandName, strings.Split(commandArgumentsString, " ")...)
				command.Stderr = os.Stderr
				command.Stdout = os.Stdout
				err = command.Run()
				if err != nil {
					continue
				}
			} else {
				err := fnc(commandArgumentsString)
				if err != nil {
				}
			}
		}
	}
}
