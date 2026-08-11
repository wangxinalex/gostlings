// Concept: time.Duration carries units and supports ordinary arithmetic.
// Task: report a duration in its original string form and whole milliseconds.
// Hint: divide by time.Millisecond only for the numeric field; keep the duration string intact.
package main

import (
	"fmt"
	"time"
)

func durationReport(duration time.Duration) string {
	// TODO: Include duration.String() and whole milliseconds in the documented format.
	return ""
}

func main() { fmt.Println(durationReport(1500 * time.Millisecond)) }
