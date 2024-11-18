package coretypes

import (
	"encoding/json"
	"fmt"
)

type UpdateType int

const (
	ProjectStabilityUpdateType UpdateType = iota
	ProjectPulseChainsUpdateType
	PulseChainUpdateType
	PulseSignatureUpdateType
	PulseUpdateType
)

const PROJECT_STABILITY_UPDATE = "ProjectStabilityUpdate"
const PROJECT_PULSE_CHAINS_UPDATE = "ProjectPulseChainsUpdate"
const PULSE_CHAIN_UPDATE = "PulseChainUpdate"
const PULSE_SIGNATURE_UPDATE = "PulseSignatureType"
const PULSE_UPDATE = "PulseUpdate"
const UNKNOWN_UPDATE = "UnknownUpdate"

func (u *UpdateType) String() string {
	switch *u {
	case ProjectStabilityUpdateType:
		return PROJECT_STABILITY_UPDATE
	case ProjectPulseChainsUpdateType:
		return PROJECT_PULSE_CHAINS_UPDATE
	case PulseChainUpdateType:
		return PULSE_CHAIN_UPDATE
	case PulseSignatureUpdateType:
		return PULSE_SIGNATURE_UPDATE
	case PulseUpdateType:
		return PULSE_UPDATE
	default:
		return UNKNOWN_UPDATE
	}
}

func ParseUpdateType(updateType string) (UpdateType, error) {
	switch updateType {
	case PROJECT_STABILITY_UPDATE:
		return ProjectStabilityUpdateType, nil
	case PROJECT_PULSE_CHAINS_UPDATE:
		return ProjectPulseChainsUpdateType, nil
	case PULSE_CHAIN_UPDATE:
		return PulseChainUpdateType, nil
	case PULSE_SIGNATURE_UPDATE:
		return PulseSignatureUpdateType, nil
	default:
		return -1, fmt.Errorf("invalid status: %s", updateType)
	}
}

func (u *UpdateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *UpdateType) UnmarshalJSON(data []byte) error {
	var updateTypeString string
	if err := json.Unmarshal(data, &updateTypeString); err != nil {
		return err
	}

	parsedUpdateType, err := ParseUpdateType(updateTypeString)
	if err != nil {
		return err
	}

	*u = parsedUpdateType
	return nil
}
