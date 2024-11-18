package utils

import "github.com/google/uuid"

// Generates a new UUID and returns it as a string
func GenerateUUID() string {
	return uuid.New().String()
}
