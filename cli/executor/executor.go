package executor

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Executor interface {
	ExecuteCommand(command string, args ...string) (string, error)
}

// Executor is a wrapper around the CLI tool for executing commands.
type executor struct {
}

// NewExecutor creates a new Executor instance.
func NewExecutor() (Executor, error) {
	return &executor{}, nil
}

// ExecuteCommand runs a CLI command with the provided arguments and returns the result.
func (e *executor) ExecuteCommand(command string, args ...string) (string, error) {
	// Split command tokens (e.g., "send event" becomes ["send", "event"])
	commandTokens := strings.Split(command, " ")
	cmdArgs := append(commandTokens, args...)

	// Create the exec.Command
	cluttaCmd := "clutta"
	cmd := exec.Command(cluttaCmd, cmdArgs...)

	// Capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute command: %s, error: %v", stderr.String(), err)
	}

	// Return trimmed stdout
	return strings.TrimSpace(stdout.String()), nil
}
