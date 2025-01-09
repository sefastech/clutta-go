package executor

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockExecutor provides a mock implementation for testing.
type MockExecutor struct {
	cliPath string
}

func (m *MockExecutor) ExecuteCommand(command string, args ...string) (string, error) {
	if m.cliPath == "" {
		return "", errors.New("CLI path is not set")
	}
	if command == "" {
		return "", errors.New("command cannot be empty")
	}
	return "mock output", nil
}

func TestNewExecutor_Success(t *testing.T) {
	cliPath := "/usr/local/bin/clutta"
	exec, err := NewExecutor(cliPath)

	assert.NoError(t, err)
	assert.NotNil(t, exec)
}

func TestNewExecutor_EmptyCLIPath(t *testing.T) {
	cliPath := ""
	exec, err := NewExecutor(cliPath)

	assert.Error(t, err)
	assert.Nil(t, exec)
	assert.Equal(t, "cliPath cannot be empty", err.Error())
}

func TestExecuteCommand_Success(t *testing.T) {
	// Mocked executor
	mockExec := &MockExecutor{cliPath: "/usr/local/bin/clutta"}

	output, err := mockExec.ExecuteCommand("send event", "--json", `{"key":"value"}`)
	assert.NoError(t, err)
	assert.Equal(t, "mock output", output)
}

func TestExecuteCommand_Failure(t *testing.T) {
	// Mocked executor
	mockExec := &MockExecutor{cliPath: ""}

	output, err := mockExec.ExecuteCommand("send event", "--json", `{"key":"value"}`)
	assert.Error(t, err)
	assert.Empty(t, output)
	assert.Equal(t, "CLI path is not set", err.Error())
}

func TestExecuteCommand_EmptyCommand(t *testing.T) {
	// Mocked executor
	mockExec := &MockExecutor{cliPath: "/usr/local/bin/clutta"}

	output, err := mockExec.ExecuteCommand("")
	assert.Error(t, err)
	assert.Empty(t, output)
	assert.Equal(t, "command cannot be empty", err.Error())
}
