package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/executor"
	"github.com/codecrafters-io/shell-starter-go/parser"
	"github.com/codecrafters-io/shell-starter-go/processor"
)

func main() {
	for {
		fmt.Print("$ ")
		input := bufio.NewScanner(bufio.NewReader(os.Stdin))
		if input.Scan() {
			userCommand := input.Text()
			newParser := parser.NewParser()
			newParser.Parse(userCommand)
			arguments := newParser.Args()
			command := processor.ProcessCommandArgs(arguments)
			executor.ExecuteCommand(command)
		}
	}
}
