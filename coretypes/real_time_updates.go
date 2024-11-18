package coretypes

import "time"

type ProjectStabilityUpdate struct {
	UUID            string    `json:"uuid" db:"uuid" gorm:"type:uuid;primaryKey"`
	ProjectID       string    `json:"projectId" db:"project_id"`
	SuccessCount    int64     `json:"successCount" db:"success_count"`
	FailureCount    int64     `json:"failureCount" db:"failure_count"`
	TotalEventCount int64     `json:"totalEventCount" db:"total_event_count"`
	UptimeInMs      float64   `json:"uptimeInHours" db:"uptime_in_ms"`
	DowntimeInMs    float64   `json:"downtimeInMs" db:"downtime_in_ms"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
}

type ProjectPulseChainsUpdate struct {
	UUID                        string    `json:"uuid" db:"uuid" gorm:"type:uuid;primaryKey"`
	ProjectID                   string    `json:"projectId" db:"project_id"`
	PulseChainCount             int       `json:"pulseChainCount" db:"pulse_chain_count"`
	ProjectErrorCount           int64     `json:"projectErrorCount" db:"project_error_count"`
	ProjectSuccessCount         int64     `json:"projectSuccessCount" db:"project_success_count"`
	ProjectPulseCount           int64     `json:"projectPulseCount" db:"project_pulse_count"`
	ProjectCompletedPulsesCount int64     `json:"projectCompletedPulsesCount" db:"project_completed_pulses_count"`
	CreatedAt                   time.Time `json:"createdAt" db:"created_at"`
}

type PulseChainUpdate struct {
	UUID                 string    `json:"uuid" db:"uuid" gorm:"type:uuid;primaryKey"`
	PulseChainID         string    `json:"pulseChainId" db:"pulse_chain_id"`
	SignatureCount       int64     `json:"signatureCount" db:"signature_count"`
	ActiveSignatureCount int64     `json:"activeSignatureCount" db:"active_signature_count"`
	PulseCount           int64     `json:"pulseCount" db:"pulse_count"`
	ErrorPulseCount      int64     `json:"errorPulseCount" db:"error_pulse_count"`
	SuccessPulseCount    int64     `json:"successPulseCount" db:"success_pulse_count"`
	CompletedPulsesCount int64     `json:"completedPulsesCount" db:"completed_pulses_count"`
	CreatedAt            time.Time `json:"createdAt" db:"created_at"`
}

type PulseSignatureUpdate struct {
	UUID                 string    `json:"uuid" db:"uuid" gorm:"type:uuid;primaryKey"`
	PulseSignatureID     string    `json:"pulseSignatureId" db:"pulse_signature_id"`
	PulseCount           int64     `json:"pulseCount" db:"pulse_count"`
	ErrorPulseCount      int64     `json:"errorPulseCount" db:"error_pulse_count"`
	SuccessPulseCount    int64     `json:"successPulseCount" db:"success_pulse_count"`
	CompletedPulsesCount int64     `json:"completedPulsesCount" db:"completed_pulses_count"`
	CreatedAt            time.Time `json:"createdAt" db:"created_at"`
}

type PulseUpdate struct {
	UUID             string    `json:"uuid" db:"uuid" gorm:"type:uuid;primaryKey"`
	PulseID          string    `json:"pulseId" db:"pulse_id"`
	Source           string    `json:"source" db:"source"`
	SourcePulseCount int64     `json:"sourcePulseCount" db:"source_pulse_count"`
	SentAt           time.Time `json:"sentAt" db:"sent_at"`
	ReceivedAt       time.Time `json:"receivedAt" db:"received_at"`
	HasIssues        bool      `json:"hasIssues" db:"has_issues"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
}

type PulseSignatureIssuesUpdate struct {
	UUID              string `json:"uuid" db:"uuid" gorm:"type:uuid;primaryKey"`
	PulseSignatureID  string `json:"pulseSignatureId" db:"pulse_signature_id"`
	PulseID           string `json:"pulseId" db:"pulse_id"`
	InvalidCreatedAt  bool   `json:"invalidCreatedAt" db:"invalid_created_at"`
	InvalidReceivedAt bool   `json:"invalidReceivedAt" db:"invalid_received_at"`
	TimelyPulse       bool   `json:"timelyPulse" db:"timely_pulse"`
	InvalidPulse      bool   `json:"invalidPulse" db:"invalid_pulse"`
}

type PulseChainIssuesUpdate struct {
	UUID                   string `json:"uuid" db:"uuid" gorm:"type:uuid;primaryKey"`
	PulseChainID           string `json:"pulseChainId" db:"pulse_chain_id"`
	PulseID                string `json:"pulseId" db:"pulse_id"`
	NewPulseIsInOrder      bool   `json:"newPulseIsInOrder" db:"new_pulse_is_in_order"`
	ChainCompletedForPulse bool   `json:"chainCompletedForPulse" db:"chain_completed_for_pulse"`
	AllPulsesArrivedInTime bool   `json:"allPulsesArrivedInTime" db:"all_pulses_arrived_in_time"`
}

type ProjectIssuesUpdate struct {
	UUID      string `json:"uuid" db:"uuid" gorm:"type:uuid;primaryKey"`
	ProjectID string `json:"projectId" db:"project_id"`
}
