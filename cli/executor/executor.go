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
	cliPath string // Path to the CLI binary
}

// NewExecutor creates a new Executor instance.
func NewExecutor(cliPath string) (Executor, error) {
	// Ensure the CLI path exists
	if cliPath == "" {
		return nil, fmt.Errorf("cliPath cannot be empty")
	}
	return &executor{cliPath: cliPath}, nil
}

// ExecuteCommand runs a CLI command with the provided arguments and returns the result.
func (e *executor) ExecuteCommand(command string, args ...string) (string, error) {
	if e.cliPath == "" {
		return "", fmt.Errorf("CLI path is not set")
	}

	// Split command tokens (e.g., "send event" becomes ["send", "event"])
	commandTokens := strings.Split(command, " ")
	cmdArgs := append(commandTokens, args...)

	// Create the exec.Command
	cmd := exec.Command(e.cliPath, cmdArgs...)

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
