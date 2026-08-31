package executor

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd"
)

func ExecuteCommand(command *cmd.Command) {
	context := &cmd.ExecutionContext{Stderr: os.Stderr, Stdout: os.Stdout}
	if command.StdoutFile != "" {
		file, err := os.OpenFile(command.StdoutFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		defer file.Close()
		context.Stdout = file
	}
	if command.StderrFile != "" {
		file, err := os.OpenFile(command.StderrFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		defer file.Close()
		context.Stderr = file
	}
	executable, ok := cmd.BuiltinsMap[command.Name]
	if !ok {
		externalCommand := &cmd.ExternalCommandExecutor{}
		err := externalCommand.Execute(command, context)
		if err != nil {
			return
		}
		return
	}
	err := executable.Execute(command, context)
	if err != nil {
		return
	}
}
