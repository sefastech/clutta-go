package clutta_sync

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	pb "github.com/sefastech/clutta-go/grpc/clutta-sync"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client interface {
	SendPulse(chainId string,
		signatureId string,
		correlationId string,
		sourceId string,
		status string,
		status_description string) error
}

type syncClient struct {
	client            pb.PulseStreamServiceClient
	conn              *grpc.ClientConn
	subscriptionValid bool
}


func SetupClient(customerId string, host string, port string) (Client, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//TLS must be enabled
	creds := credentials.NewTLS(nil)

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %s", err)
	}

	client := pb.NewPulseStreamServiceClient(conn)

	// Check customer's subscription status
	valid, err := checkCustomerSubscription(customerId)
	if err != nil {
		closeErr := conn.Close() // Ensure the connection is closed if subscription check fails
		handleCloseConnectionError(closeErr)
		return nil, fmt.Errorf("subscription check failed: %s", err)
	}

	return &syncClient{client: client, conn: conn, subscriptionValid: valid}, nil
}

// 2. Function to check customer subscription validity
func checkCustomerSubscription(customerId string) (bool, error) {
	// Placeholder for HTTP call or database query to check subscription status
	// Example: HTTP GET to https://api.example.com/subscriptions?customerId={customerId}
	return true, nil
}


func (c *syncClient) SendPulse(chainId string,
	signatureId string,
	correlationId string,
	sourceId string,
	status string,
	statusDescription string) error {

	if !c.subscriptionValid {
		return fmt.Errorf("client does not have a valid subscription to send events")
	}

	pulseStatusRes, exists := pb.PulseStatus_value[status]
	var pulseStatus pb.PulseStatus

	if !exists {
		pulseStatus = pb.PulseStatus_UNKNOWN
	} else {
		pulseStatus = pb.PulseStatus(pulseStatusRes)
	}

	pulse := pb.Pulse{
		Uuid:              uuid.NewString(),
		ChainId:           chainId,
		SignatureId:       signatureId,
		CorrelationId:     correlationId,
		SourceId:          sourceId,
		Status:            pulseStatus,
		StatusDescription: statusDescription,
		CreatedAt:         time.Now().UnixMilli(),
	}

	stream, clientErr := c.client.SendPulse(context.Background())
	if clientErr != nil {
		return fmt.Errorf("error creating stream: %v", clientErr)
	}

	if streamErr := stream.Send(&pulse); streamErr != nil {
		return fmt.Errorf("failed to send pulse: %v", streamErr)
	}

	fmt.Println("Pulse sent successfully!")

	_,closeErr := stream.CloseAndRecv()
	if closeErr != nil {
		return fmt.Errorf("failed to close stream: %v", closeErr)
	}

	Logger.Infof("Sent pulse with UUID: %s", pulse.Uuid)
	return nil
}
