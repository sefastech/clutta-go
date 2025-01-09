package installer

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// getInstalledVersion retrieves the version of the installed CLI tool.
func getInstalledVersion(cliPath string, ops SystemOps) (string, error) {
	cmd := ops.Command(cliPath, "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get installed version: %w", err)
	}

	versionString := strings.TrimSpace(string(output))
	return extractVersion(versionString)
}

// extractVersion extracts the version string from the input text.
// Returns the version string if found, otherwise returns an empty string.
func extractVersion(input string) (string, error) {
	// Regular expression to match the version
	re := regexp.MustCompile(`Version:\s*(v[\d.]+)`)

	// Find the match
	matches := re.FindStringSubmatch(input)

	// Return the captured version if it exists
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("version not found in input: %s", input)
}

// getLatestVersion fetches the latest release version from GitHub.
func getLatestVersion(repo string, ops SystemOps) (string, error) {
	release, err := fetchLatestRelease(repo, ops)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest release: %w", err)
	}
	return release.TagName, nil
}

// checkCLI verifies if the CLI tool is installed and executable.
func checkCLI(cliPath string, ops SystemOps) bool {
	// Check if the file exists
	if _, err := ops.Stat(cliPath); os.IsNotExist(err) {
		return false
	}

	// Ensure the file is executable (not required for Windows)
	if runtime.GOOS != "windows" {
		if err := ops.Chmod(cliPath, 0755); err != nil {
			fmt.Printf("Warning: CLI tool at %s is not executable.\n", cliPath)
			return false
		}
	}

	return true
}

// ensureDependencies checks if required tools (e.g., curl) are available.
func ensureDependencies(ops SystemOps) error {
	// Check for curl
	if _, err := ops.LookPath("curl"); err != nil {
		return fmt.Errorf("curl is required but not installed. Please install curl and try again")
	}
	return nil
}

// isRunningInContainer detects if the code is running inside a container.
func isRunningInContainer(ops SystemOps) bool {
	if _, err := ops.Stat("/.dockerenv"); err == nil {
		return true // Docker container
	}
	if content, err := ops.ReadFile("/proc/1/cgroup"); err == nil {
		return strings.Contains(string(content), "docker") || strings.Contains(string(content), "kubepods")
	}
	return false
}
