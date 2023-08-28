package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type adapter struct{}

func (a *adapter) Write(data []byte) (int, error) {
	fmt.Printf("WRITE <%s>\n", string(data))
	return len(data), nil
}

func main() {
	cmd := exec.Command("./command.sh")
	// When using the adapter, cmd.Run() waits for the nohup command to exit.
	cmd.Stdout = &adapter{}
	// When using Stdout, a *File, cmd.Run() returns before nohup completes.
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			fmt.Printf("commanded exited with %d\n", exitErr.ExitCode())
		} else {
			fmt.Printf("failed to run command: %s\n", err)
		}
	}
}
