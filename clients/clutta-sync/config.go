package clutta_sync

import "github.com/sefastech/clutta/libraries/golang/logging"

var Logger = logging.InitializeLogger("info")

func SetLogger(logger logging.Logger) {
	Logger = logger
}
