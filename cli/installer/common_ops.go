package installer

import (
	"net/http"
	"os"
	"os/exec"
)

// SystemOps defines an interface for system-level operations, enabling easy mocking during tests.
type SystemOps interface {
	Stat(path string) (os.FileInfo, error)
	Chmod(path string, mode os.FileMode) error
	LookPath(file string) (string, error)
	ReadFile(path string) ([]byte, error)
	Command(name string, arg ...string) CommandRunner
	FetchFromURL(url string) (resp *http.Response, err error)
}

// CommandRunner defines an interface for executing commands.
type CommandRunner interface {
	Output() ([]byte, error)
}

// defaultSystemOps provides the default implementation of SystemOps using standard library functions.
type defaultSystemOps struct{}

func NewSystemOps() SystemOps {
	return &defaultSystemOps{}
}
func (d *defaultSystemOps) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func (d *defaultSystemOps) Chmod(path string, mode os.FileMode) error {
	return os.Chmod(path, mode)
}

func (d *defaultSystemOps) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

func (d *defaultSystemOps) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (d *defaultSystemOps) Command(name string, arg ...string) CommandRunner {
	return exec.Command(name, arg...)
}

func (d *defaultSystemOps) FetchFromURL(url string) (resp *http.Response, err error) {
	return http.Get(url)
}
