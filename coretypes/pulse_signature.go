package coretypes

import "time"

type PulseSignature struct {
	UUID                   string    `json:"uuid"`
	Name                   string    `json:"name"`
	ChainId                string    `json:"chainId"`
	ParentSignatureId      string    `json:"parentSignatureId"`
	IsRootSignature        bool      `json:"isRootSignature"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
	SourceIdentifierPrefix string    `json:"sourceIdentifierPrefix"`
	PriorityLevel          string    `json:"priorityLevel"`
	ArrivalTimeInMs        int64     `json:"arrivalTimeInMs"`
	Description            string    `json:"description"`
}
