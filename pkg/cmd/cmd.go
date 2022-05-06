package cmd

import (
	"os"
	"os/exec"
)

func ExecuteCommand(command string) error {
	cmd := exec.Command(os.Getenv("windir")+`\system32\cmd.exe`, "/C", command)
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}

func ExecuteCommandWithOutput(command string) (string, error) {
	cmd := exec.Command(os.Getenv("windir")+`\system32\cmd.exe`, "/C", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
