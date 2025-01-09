package info

import (
	"fmt"
	"runtime"
)

// OS Constants for installation paths
const (
	InstallPathWindows = `C:\Program Files\clutta`
	InstallPathLinux   = "/usr/local/bin"
	InstallPathDarwin  = "/usr/local/bin"
)

// CLIInfo holds metadata about the CLI tool.
type CLIInfo struct {
	Name        string // CLI tool name (e.g., "my-cli-tool")
	Repo        string // GitHub repository in "owner/repo" format
	BinaryName  string // Expected binary name after installation
	InstallPath string // Path where the binary will be installed
}

// DefaultCLIInfo returns a preconfigured CLIInfo based on the operating system.
func DefaultCLIInfo(osType string) (CLIInfo, error) {
	var installPath string
	switch osType {
	case "windows":
		installPath = InstallPathWindows
	case "linux":
		installPath = InstallPathLinux
	case "darwin":
		installPath = InstallPathDarwin
	default:
		return CLIInfo{}, fmt.Errorf("unsupported operating system: %s", osType)
	}

	return CLIInfo{
		Name:        "clutta-cli",                    // Name for downloading
		BinaryName:  "clutta" + getBinaryExtension(), // Name for the installed binary
		Repo:        "sefastech/clutta-cli-releases", // GitHub repository
		InstallPath: installPath,                     // Install path for Windows/Linux/macOS
	}, nil
}

// getBinaryExtension returns the appropriate binary extension based on the OS.
func getBinaryExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}
