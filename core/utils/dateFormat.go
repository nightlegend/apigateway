package utils

import (
	"time"
)

func GetCurrentTime() *time.Time {
	t := time.Now()
	return &t
}

func GetCurrentTimeByYYYY_MM_DD() string {
	t := time.Now()
	currentTime := t.Format("2006-Jan-02")
	return currentTime
}

func GetCurrentTimeByYYYY_MM_DDHH_MM_SS() string {
	t := time.Now()
	currentTime := t.Format("2006-Jan-02 15:04:05")
	return currentTime
}
