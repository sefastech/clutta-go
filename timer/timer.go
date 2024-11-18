package timer

import (
	"log"
	"strconv"
	"time"
)

func TimeElapsedMs(cummulativeElapsedTime float64, startTime time.Time) float64 {
	elapsedTimeDuration := time.Since(startTime).Milliseconds()
	return cummulativeElapsedTime + float64(elapsedTimeDuration)
}

func TimeElapsedMsWithEndTime(cummulativeElapsedTime float64, startTime, endTime time.Time) float64 {
	elapsedTimeDuration := endTime.Sub(startTime).Milliseconds()
	return cummulativeElapsedTime + float64(elapsedTimeDuration)
}

func TimeElapsedMsSinceCreatedUnix(createdAtUnix int64) int64 {
	createdAt := time.Unix(createdAtUnix, 0)
	now := time.Now()
	duration := now.Sub(createdAt)
	return duration.Milliseconds()
}
func TimeElapsedMsSinceCreatedUnixWithNow(createdAtUnix int64, now time.Time) int64 {
	createdAt := time.Unix(createdAtUnix, 0)
	duration := now.Sub(createdAt)
	return duration.Milliseconds()
}

func ParseUnixTimestamp(createdAt string) int64 {
	timestamp, err := strconv.ParseInt(createdAt, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse timestamp: %v", err)
	}
	return timestamp
}

func ParseStringTimestamp(createdAt int64) string {
	timestamp := int64(createdAt) // Unix timestamp
	formattedTime := time.Unix(timestamp, 0).UTC().Format("02-01-2006 15:04:05")
	return formattedTime
}
