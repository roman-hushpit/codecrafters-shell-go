package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Print("$ ")
		input := bufio.NewScanner(bufio.NewReader(os.Stdin))
		if input.Scan() {
			userInput := input.Text()
			fmt.Printf("%s: command not found\n", userInput)
		}
	}
}
