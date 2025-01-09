# Clutta-Go SDK

Clutta-Go is the official Go SDK for interacting with Clutta, a platform for observability and monitoring. 
Use this SDK to send pulse data, either individually or in batches, with ease and reliability.

---

## Installation

To add Clutta-Go to your Go project, use the following command:

```go
go get github.com/sefastech/clutta-go
```

## Environment Variables
`CLUTTA_AUTO_UPDATE`

The `CLUTTA_AUTO_UPDATE` environment variable enables automatic downloading of the latest Clutta CLI version. 
When set to true, the SDK will automatically check for and download updates to ensure you're always running the latest version of the CLI.

Example usage:
```bash
export CLUTTA_AUTO_UPDATE=true
```
> Note: Enabling this feature ensures you stay up to date with the latest improvements, bug fixes, 
> and enhancements. However, you can disable it by setting the variable to false or not setting it at all.

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
	"yourorg/clutta-go/lib"
)

client, err := lib.NewClient()
if err != nil {
	log.Fatalf("Failed to create client: %v", err)
}
fmt.Println("Clutta-Go client initialized successfully.")

```

## 2. Sending a Single Pulse
   Send an individual pulse for targeted updates.
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
Optimize performance by sending pulses in bulk
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