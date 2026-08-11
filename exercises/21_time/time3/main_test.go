package main

import (
	"testing"
	"time"
)

func TestDurationReportPreservesUnitsAndBoundaries(t *testing.T) {
	tests := []struct {
		name string
		in   time.Duration
		want string
	}{
		{"fractional second", 1500 * time.Millisecond, "duration=1.5s millis=1500"},
		{"sub millisecond", 750 * time.Microsecond, "duration=750µs millis=0"},
		{"negative", -2 * time.Second, "duration=-2s millis=-2000"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := durationReport(test.in); got != test.want {
				t.Fatalf("durationReport(%v) = %q, want %q", test.in, got, test.want)
			}
		})
	}
}
