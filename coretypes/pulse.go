package coretypes

import "time"

// PulseStatus is an enumeration of possible pulse statuses.
type PulseStatus int

const (
	Unknown PulseStatus = iota // Default to 0
	Success
	Failure
)

type Pulse struct {
	UUID              string      `json:"uuid"`
	ChainId           string      `json:"chainId"`
	SignatureId       string      `json:"signatureId"`
	CorrelationId     string      `json:"correlationId"`
	SourceId          string      `json:"sourceId"`
	Status            PulseStatus `json:"pulseStatus"`
	StatusDescription string      `json:"statusDescription"`
	CreatedAt         time.Time   `json:"createdAt"`
	ReceivedAt        time.Time   `json:"receivedAt"`
}
