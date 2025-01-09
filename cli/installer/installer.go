package installer

import (
	"fmt"
	"github.com/sefastech/clutta-go/cli/info"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// EnsureCLI ensures that the CLI tool is installed and up to date.
func EnsureCLI(info info.CLIInfo) error {
	cliPath := filepath.Join(info.InstallPath, info.BinaryName)
	systemOps := NewSystemOps()

	// Check if CLI is installed
	if !checkCLI(cliPath, systemOps) {
		fmt.Println("CLI tool not found. Installing the latest version...")
		return installLatestCLI(info, systemOps)
	}

	// Check if CLI_AUTO_UPDATE is enabled
	autoUpdate := os.Getenv("CLUTTA_AUTO_UPDATE")
	if strings.ToLower(autoUpdate) != "true" {
		fmt.Println("CLUTTA_AUTO_UPDATE is not enabled. Skipping update check.")
		return nil
	}

	// Fetch the current and latest versions
	currentVersion, err := getInstalledVersion(cliPath, systemOps)
	if err != nil {
		fmt.Printf("Warning: Could not determine installed version: %v\n", err)
		return nil
	}

	latestVersion, err := getLatestVersion(info.Repo, systemOps)
	if err != nil {
		fmt.Printf("Warning: Could not determine the latest version: %v\n", err)
		return nil
	}

	// Compare versions and update if stale
	if currentVersion != latestVersion {
		fmt.Printf("CLI tool is outdated (current: %s, latest: %s). Updating...\n", currentVersion, latestVersion)
		return installLatestCLI(info, systemOps)
	}

	fmt.Println("CLI tool is up to date.")
	return nil
}

// installLatestCLI fetches and installs the latest version of the CLI tool
func installLatestCLI(info info.CLIInfo, ops SystemOps) error {
	// Step 1: Ensure dependencies are installed
	if err := ensureDependencies(ops); err != nil {
		return err
	}

	// Step 2: Determine the appropriate download URL
	downloadURL, err := getDownloadURL(info)
	if err != nil {
		return fmt.Errorf("failed to determine download URL: %w", err)
	}

	// Step 3: Ensure the install directory exists
	if err := os.MkdirAll(info.InstallPath, 0755); err != nil {
		return fmt.Errorf("failed to create install directory: %w", err)
	}

	// Step 4: Download the binary to the install path
	finalPath, err := downloadBinary(downloadURL, info.InstallPath, info.BinaryName, ops)
	if err != nil {
		return fmt.Errorf("failed to download and install binary: %w", err)
	}

	// Step 5: Update PATH for Windows
	if runtime.GOOS == "windows" {
		if err := addToWindowsPath(info.InstallPath); err != nil {
			return fmt.Errorf("failed to update PATH: %w", err)
		}
	}

	// Step 6: Make the binary executable (Linux/macOS only)
	if runtime.GOOS != "windows" {
		if err := os.Chmod(finalPath, 0755); err != nil {
			return fmt.Errorf("failed to set executable permissions: %w", err)
		}
	}

	fmt.Printf("CLI tool installed successfully at %s\n", finalPath)
	return nil
}
