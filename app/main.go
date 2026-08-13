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
			fnc, ok := cmd.BuiltinsMap[arguments[0]]
			if !ok {
				executable, err := cmd.FindExecutable(arguments[0])
				if err != nil {
					fmt.Printf("%s: command not found\n", command)
					continue
				}
				command := exec.Command(executable, arguments[1:]...)
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
