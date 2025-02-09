package checker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

// EnsureCLI ensures that the CLI tool is installed and up to date.
func EnsureCLI() error {
	// 1. supported OS?
	if !supportedOS(runtime.GOOS) {
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// 2. clutta is not installed
	if !CluttaIsInstalled() {
		return fmt.Errorf("Clutta is NOT installed")
	}

	// 3a. get installed version
	installedVersion, getErr := getInstalledVersion()
	if getErr != nil {
		return getErr
	}

	// 3b. get latest version
	latestVersion, getErr := getLatestVersion()
	if getErr != nil {
		return getErr
	}

	// 4. get installed version
	if installedVersion != latestVersion {
		fmt.Printf("Clutta is outdated! Installed: %s, Latest: %s\n", installedVersion, latestVersion)
		return nil
	}

	// 5. all good
	fmt.Println("CLI tool is up to date.")
	return nil
}

// CluttaIsInstalled checks if the "clutta" command is available and executable.
func CluttaIsInstalled() bool {
	cmd := "clutta"

	// 1. explicit check command
	checkCmd := "which"
	if runtime.GOOS == "windows" {
		checkCmd = "where"
	}

	// 2. run command
	if _, err := exec.LookPath(cmd); err != nil {
		// If not found, run an explicit check using "which" or "where"
		if _, err := exec.Command(checkCmd, cmd).Output(); err != nil {
			return false
		}
	}

	// 3. check if it runs successfully
	if err := exec.Command(cmd, "--version").Run(); err != nil {
		return false
	}

	return true
}

// getInstalledVersion runs "clutta --version" and returns the version string.
func getInstalledVersion() (string, error) {
	out, err := exec.Command("clutta", "--version").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// getLatestVersion fetches the latest release from GitHub.
func getLatestVersion() (string, error) {
	resp, err := http.Get("https://api.github.com/repos/sefastech/clutta-cli-releases/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		TagName string `json:"tag_name"`
	}
	if decodeErr := json.NewDecoder(resp.Body).Decode(&result); decodeErr != nil {
		return "", decodeErr
	}
	return result.TagName, nil
}

func supportedOS(ostype string) bool {
	return ostype == "windows" || ostype == "darwin" || ostype == "linux"
}
