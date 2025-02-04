package installer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// addToWindowsPath adds a directory to the PATH environment variable on Windows if not already present.
func addToWindowsPath(dir string) error {
	// Check if the directory is already in PATH
	currentPath := os.Getenv("PATH")
	if strings.Contains(currentPath, dir) {
		fmt.Printf("Directory %s is already in PATH. Skipping update.\n", dir)
		return nil
	}

	fmt.Printf("Adding %s to PATH...\n", dir)

	// PowerShell command to update the PATH environment variable
	command := fmt.Sprintf(`[Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";%s", "User")`, dir)
	cmd := exec.Command("powershell", "-Command", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update PATH: %w", err)
	}

	fmt.Println("Successfully added CLI tool to PATH.")
	return nil
}

// moveToInstallPath moves the binary to the final install path on all platforms.
func moveToInstallPath(source, destination string, ops SystemOps) error {
	if runtime.GOOS == "windows" {
		// Check if the destination requires admin privileges
		if strings.HasPrefix(destination, `C:\Program Files`) {
			fmt.Println("Warning: Installing to a protected path requires admin privileges.")
		}

		// Use os.Rename for simple moves
		return os.Rename(source, destination)
	}

	// For Linux/macOS, use sudo if necessary
	if isRunningInContainer(ops) {
		return os.Rename(source, destination)
	}

	cmd := getMoveBinaryCmd(source, destination)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to move binary with sudo: %w", err)
	}

	return nil
}

func getMoveBinaryCmd(source, destination string) *exec.Cmd {
	// Check if sudo is available
	_, err := exec.LookPath("sudo")
	useSudo := err == nil // sudo exists

	// Prepare command
	cmd := exec.Command("mv", source, destination)
	if useSudo {
		cmd = exec.Command("sudo", "mv", source, destination)
	}
	return cmd
}
