// Concept: Go's unique date/time layout — the reference time Mon Jan 2 15:04:05 MST 2006
// Task: fill in the layout string so Format prints the time as shown
// Expected output: Format: 2026-01-02 15:04:05
// Hint: Go layout strings use the reference time itself, not %Y-style verbs (Go doc: time)

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
