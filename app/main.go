package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd"
)

func main() {
	for {
		fmt.Print("$ ")
		input := bufio.NewScanner(bufio.NewReader(os.Stdin))
		if input.Scan() {
			command := input.Text()
			fnc, ok := cmd.BuiltinsMap[command]
			if !ok {
				fmt.Printf("%s: command not found\n", command)
				continue
			}

			err := fnc()
			if err != nil {
			}
		}
	}
}
