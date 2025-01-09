package installer

import (
	"encoding/json"
	"fmt"
	"github.com/sefastech/clutta-go/cli/info"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// getDownloadURL determines the appropriate download URL for the CLI binary.
func getDownloadURL(info info.CLIInfo) (string, error) {
	arch := runtime.GOARCH // Architecture (e.g., "amd64", "arm64")
	os := runtime.GOOS     // Operating system (e.g., "linux", "darwin", "windows")

	// Construct the binary filename based on OS and architecture
	var filename string
	switch os {
	case "linux":
		filename = fmt.Sprintf("%s_linux_%s", info.Name, arch)
	case "darwin":
		filename = fmt.Sprintf("%s_darwin_%s", info.Name, arch)
	case "windows":
		filename = fmt.Sprintf("%s_windows_%s.exe", info.Name, arch)
	default:
		return "", fmt.Errorf("unsupported operating system: %s", os)
	}

	// Construct the download URL
	return fmt.Sprintf("https://github.com/%s/releases/latest/download/%s", info.Repo, filename), nil
}

// downloadBinary downloads the binary from the URL to the specified install path.
func downloadBinary(url, installPath, binaryName string, ops SystemOps) (string, error) {
	tempDir := os.TempDir()                                // Temporary directory
	tempFile := filepath.Join(tempDir, filepath.Base(url)) // Temporary file path

	fmt.Printf("Downloading CLI tool to temporary location: %s\n", tempFile)

	// Download the file to the temporary location
	cmd := exec.Command("curl", "-L", "-o", tempFile, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to download binary: %w", err)
	}

	// Final path for the binary
	finalPath := filepath.Join(installPath, binaryName)

	// Move and rename the binary to the install path
	fmt.Printf("Moving CLI tool to final location: %s\n", finalPath)

	if err := moveToInstallPath(tempFile, finalPath, ops); err != nil {
		return "", fmt.Errorf("failed to move binary to install path: %w", err)
	}

	// Return the final path for the installed binary
	return finalPath, nil
}

// Release represents the structure of a GitHub release.
type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

// fetchLatestRelease retrieves the latest release information from GitHub.
func fetchLatestRelease(repo string, ops SystemOps) (*Release, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	resp, err := ops.FetchFromURL(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch latest release: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("failed to decode release JSON: %w", err)
	}

	return &release, nil
}
