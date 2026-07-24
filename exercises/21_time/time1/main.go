// Concept: Go's unique date/time format — the reference time
// Task: Format now using Go's reference time layout to get the output; the time after "Format:" must match the reference time format
// Expected output: Format: 2026-01-02 15:04:05
// (the actual date/time will differ — the format pattern is what matters)
// Hint: Go uses the magic reference time Mon Jan 2 15:04:05 MST 2006 for layout strings (Go doc: time)

package main

import (
	"fmt"
	"time"
)

func main() {
	// Go's reference time: 2006-01-02T15:04:05
	now := time.Date(2026, 1, 2, 15, 4, 5, 0, time.UTC)

	// TODO: Format `now` using the layout string so it matches the expected output.
	fmt.Println("Format:", now.Format( /* TODO */ ))
}
