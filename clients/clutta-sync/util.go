package clutta_sync

func handleCloseConnectionError(closeErr error) {
	if closeErr != nil {
		Logger.Error("Could not close the gRPC connection to clutta-sync server. Reason(s): %s", closeErr.Error())
	}
}
