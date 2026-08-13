package main

import (
	"bufio"
	"fmt"
	"os"
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
				fmt.Printf("%s: command not found\n", command)
				continue
			}

			err := fnc(arguments[1:]...)
			if err != nil {
			}
		}
	}
}
