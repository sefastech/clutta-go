package utils

import (
	"encoding/json"
	"errors"
	"github.com/sefastech/clutta-go/coretypes"
	pb "github.com/sefastech/clutta-go/grpc/clutta-sync"
	"time"
)

type PulseConverter interface {
	GrpcToCore(grpcPulse *pb.Pulse) (*coretypes.Pulse, error)
	Stringify(corePulse interface{}) (string, error)
	GrpcToCoreStatus(status pb.PulseStatus) coretypes.PulseStatus
}

type pulseConverter struct {
}

func NewPulseConverter() PulseConverter {
	return &pulseConverter{}
}

func (pc pulseConverter) GrpcToCoreStatus(status pb.PulseStatus) coretypes.PulseStatus {
	switch status {
	case pb.PulseStatus_SUCCESS:
		return coretypes.Success
	case pb.PulseStatus_FAILURE:
		return coretypes.Failure
	default:
		return coretypes.Unknown
	}
}

func (pc pulseConverter) GrpcToCore(grpcPulse *pb.Pulse) (*coretypes.Pulse, error) {
	if grpcPulse == nil {
		return nil, errors.New("could not extract system Pulse data from gRPC version")
	}

	return &coretypes.Pulse{UUID: grpcPulse.GetUuid(), SignatureId: grpcPulse.GetSignatureId(),
		CorrelationId: grpcPulse.GetCorrelationId(), SourceId: grpcPulse.GetSourceId(),
		Status: pc.GrpcToCoreStatus(grpcPulse.GetStatus()), StatusDescription: grpcPulse.GetStatusDescription(),
		CreatedAt: time.Unix(grpcPulse.CreatedAt, 0)}, nil
}

func (pc pulseConverter) Stringify(pulse interface{}) (string, error) {
	data, err := json.Marshal(pulse)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
