package lib

import (
	"encoding/json"
	"fmt"
	"github.com/sefastech/clutta-go/cli/executor"
	"github.com/sefastech/clutta-go/cli/info"
	"github.com/sefastech/clutta-go/cli/installer"
	"path/filepath"
	"runtime"
)

// Client provides a high-level API for interacting with the CLI tool.
type Client interface {
	SendPulse(pulse map[string]any) (string, error)
	SendPulses(pulses []map[string]any) (string, error)
}
type client struct {
	executor executor.Executor
}

// NewClient initializes the CLI client, ensuring the CLI tool is installed and ready to use.
func NewClient() (Client, error) {
	// Get default CLI info based on the current OS
	cliInfo, err := info.DefaultCLIInfo(runtime.GOOS)
	if err != nil {
		return nil, fmt.Errorf("unsupported operating system: %w", err)
	}

	// Ensure the CLI is installed and up to date
	if err := installer.EnsureCLI(cliInfo); err != nil {
		return nil, fmt.Errorf("failed to ensure CLI tool: %w", err)
	}

	// Initialize the executor
	cliPath := filepath.Join(cliInfo.InstallPath, cliInfo.BinaryName)
	cliExecutor, err := executor.NewExecutor(cliPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create new executor: %w", err)
	}
	return &client{executor: cliExecutor}, nil
}

// SendPulse sends a single pulse to Clutta.
func (c *client) SendPulse(pulse map[string]any) (string, error) {
	// Convert the pulse map to a JSON string
	jsonString, err := json.Marshal(pulse)
	if err != nil {
		return "", fmt.Errorf("failed to encode event as JSON: %w", err)
	}

	// Execute the `send` command with the `--json` flag
	return c.executor.ExecuteCommand("send pulse", "--json", string(jsonString))
}

// SendPulses sends multiple pulses to Clutta.
func (c *client) SendPulses(pulses []map[string]any) (string, error) {
	// Convert the event map to a JSON string
	jsonString, err := json.Marshal(pulses)
	if err != nil {
		return "", fmt.Errorf("failed to encode pulse as JSON: %w", err)
	}

	// Execute the `send` command with the `--json` flag
	return c.executor.ExecuteCommand("send pulses", "--json", string(jsonString))
}
