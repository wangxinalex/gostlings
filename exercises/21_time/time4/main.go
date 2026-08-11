// Concept: a timer must stop on every exit path.
// Task: return one timer event, or close the result when stop is requested first.
// Hint: select on timer.C and stop; defer timer.Stop in the goroutine.
package main

import "time"

func waitTimer(stop <-chan struct{}, duration time.Duration) <-chan time.Time {
	// TODO: Create a timer and close the output when cancellation wins.
	return nil
}

func main() {}
