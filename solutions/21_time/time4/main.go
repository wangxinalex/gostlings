// Concept: a timer must stop on every exit path.
// Task: return one timer event, or close the result when stop is requested first.
// Hint: select on timer.C and stop; defer timer.Stop in the goroutine.
package main

import "time"

func waitTimer(stop <-chan struct{}, duration time.Duration) <-chan time.Time {
	result := make(chan time.Time, 1)
	timer := time.NewTimer(duration)
	go func() {
		defer close(result)
		defer timer.Stop()
		select {
		case value := <-timer.C:
			result <- value
		case <-stop:
		}
	}()
	return result
}

func main() {}
