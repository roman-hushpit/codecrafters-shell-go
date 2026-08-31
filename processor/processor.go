package processor

import (
	"github.com/codecrafters-io/shell-starter-go/cmd"
)

func ProcessCommandArgs(args []string) *cmd.Command {
	command := &cmd.Command{}
	if len(args) == 1 {
		command.Name = args[0]
		return command
	}
	command.Name = args[0]
	for index := 1; index < len(args); index++ {
		if args[index] == ">" || args[index] == "1>" {
			command.StdoutFile = args[index+1]
			break
		}
		command.Args = append(command.Args, args[index])
	}

	return command
}
