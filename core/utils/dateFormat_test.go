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

func TestGetCurrentTimeByYYYY_MM_DD(t *testing.T) {
	if GetCurrentTimeByYYYY_MM_DD() != "" {
		t.Log(GetCurrentTimeByYYYY_MM_DD())
	} else {
		t.Error("error")
	}
}

func TestGetCurrentTimeByYYYY_MM_DDHH_MM_SS(t *testing.T) {
	if GetCurrentTimeByYYYY_MM_DD() != "" {
		t.Log(GetCurrentTimeByYYYY_MM_DDHH_MM_SS())
	} else {
		t.Error("format error")
	}
}
