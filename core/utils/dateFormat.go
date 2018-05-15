package utils

import (
	"time"
)

// GetCurrentTime is get currently time.
func GetCurrentTime() *time.Time {
	t := time.Now()
	return &t
}

// GetCurrentTimeByYYYYMMDD is format time by YYYY-MM-DD
func GetCurrentTimeByYYYYMMDD() string {
	t := time.Now()
	currentTime := t.Format("2006-Jan-02")
	return currentTime
}

// GetCurrentTimeByYYYYMMDDHHMMSS is format time by  YYYY-MM-DD-HH-MM-SS
func GetCurrentTimeByYYYYMMDDHHMMSS() string {
	t := time.Now()
	currentTime := t.Format("2006-Jan-02 15:04:05")
	return currentTime
}
