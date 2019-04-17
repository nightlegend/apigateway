package utils

import (
	"testing"
)

func TestGetCurrentTime(t *testing.T) {
	if GetCurrentTime() != nil {
		t.Log(GetCurrentTime())
	} else {
		t.Error()
	}
}

func TestGetCurrentTimeByYYYYMMDD(t *testing.T) {
	if GetCurrentTimeByYYYYMMDD() != "" {
		t.Log(GetCurrentTimeByYYYYMMDD())
	} else {
		t.Error("error")
	}
}

func TestGetCurrentTimeByYYYYMMDDHHMMSS(t *testing.T) {
	if GetCurrentTimeByYYYYMMDDHHMMSS() != "" {
		t.Log(GetCurrentTimeByYYYYMMDDHHMMSS())
	} else {
		t.Error("format error")
	}
}
