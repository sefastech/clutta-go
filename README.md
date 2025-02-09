# Clutta-Go SDK

Clutta-Go is the official Go SDK for interacting with Clutta, a platform for observability and monitoring. 
Use this SDK to send pulse data, either individually or in batches, with ease and reliability.

---

## Installation

To add Clutta-Go to your Go project, use the following command:

```go
go get github.com/sefastech/clutta-go
```

## Features
- **Effortless Integration**: Simple setup with reusable client initialization.
- **Single Pulse Support**: Send individual data points with precision.
- **Batch Processing**: Transmit multiple data points efficiently in one operation.
- **Automatic Updates**: Ensure you always have the latest tools with the CLUTTA_AUTO_UPDATE environment variable.

# Getting Started

## 1. Initialize the Client
   Set up the client once and reuse it across your application. 

```go
import (
	"log"
	"fmt"
	"github.com/sefastech/clutta-go/lib"
)

func main() {
	client, err := lib.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println("Clutta-Go client initialized successfully.")
}
```

## 2. Sending a Pulse
   Send an individual pulse for targeted updates. Paste the following also inside the main function above.
```go

pulse := map[string]interface{}{
	"signatureId":       "unique-signature-id",
	"chainId":           "unique-chain-id",
	"correlationId":     "unique-correlation-id",
	"sourceId":          "your-source",
	"userId":            "user-id",
	"apiKey":            "api-key",
	"status":            1, // Status: 0 for unknown, 1 for success, 2 for failure
	"statusDescription": "Success",
}

response, err := client.SendPulse(pulse)
if err != nil {
	log.Fatalf("Error sending pulse: %v", err)
}

fmt.Println("Pulse sent successfully:", response)

```

## 3. Sending Multiple Pulses
Optimize performance by sending pulses in bulk. Paste the following also inside the main function above.
```go


pulses := []map[string]interface{}{
	{
		"signatureId":       "signature-1",
		"chainId":           "chain-1",
		"correlationId":     "correlation-1",
		"sourceId":          "source-1",
		"userId":            "user-1",
		"apiKey":            "api-key",
		"status":            1,
		"statusDescription": "Operation successful",
	},
	{
		"signatureId":       "signature-2",
		"chainId":           "chain-2",
		"correlationId":     "correlation-2",
		"sourceId":          "source-2",
		"userId":            "user-2",
		"apiKey":            "api-key",
		"status":            2,
		"statusDescription": "Operation failed",
	},
}

response, err := client.SendPulses(pulses)
if err != nil {
	log.Fatalf("Error sending pulses: %v", err)
}

fmt.Println("Pulses sent successfully:", response)

```

# License

Clutta-Go is open-source and licensed under the MIT License. Contributions are welcome.


# Support

For technical support, documentation, or to report issues, visit the Clutta Documentation or contact our support team.


# About Clutta

Clutta is redefining observability with cutting-edge tools designed to help organizations monitor, analyze, and optimize their systems. 
Whether you're scaling to millions of users or managing critical infrastructure, Clutta provides the insights you need to excel.
