package cmd

import (
	"os"
	"os/exec"
)

func ExecuteCommand(command string, args ...string) error {
	winDir := os.Getenv("windir")

	commandArguments := []string{"/C"}
	commandArguments = append(commandArguments, command)
	commandArguments = append(commandArguments, args...)

	cmd := exec.Command(winDir+`\system32\cmd.exe`, commandArguments...)
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}

func ExecuteCommandWithOutput(command string, args ...string) (string, error) {
	winDir := os.Getenv("windir")

	commandArguments := []string{"/C"}
	commandArguments = append(commandArguments, command)
	commandArguments = append(commandArguments, args...)

	cmd := exec.Command(winDir+`\system32\cmd.exe`, commandArguments...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
