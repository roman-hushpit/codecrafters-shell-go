package executor

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd"
)

func ExecuteCommand(command *cmd.Command) {
	context := &cmd.ExecutionContext{Stderr: os.Stderr, Stdout: os.Stdout}
	if command.StdoutFile != "" {
		file, err := openRedirectFile(command.StdoutFile, command.AppendOut)
		if err != nil {
			return
		}
		defer file.Close()
		context.Stdout = file
	}

	if command.StderrFile != "" {
		file, err := openRedirectFile(command.StderrFile, command.AppendErr)
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

func openRedirectFile(path string, appendMode bool) (*os.File, error) {
	flags := os.O_CREATE | os.O_WRONLY

	if appendMode {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}

	return os.OpenFile(path, flags, 0666)
}
