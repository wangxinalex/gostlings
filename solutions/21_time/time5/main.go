// Concept: select can distinguish a result, a closed result stream, and a timeout.
// Task: return result, canceled, or timed out without leaking the timer.
// Hint: use a timer and check the receive's ok value for the cancellation branch.
package main

import "time"

func awaitOrCancel(result <-chan string, timeout time.Duration) string {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case value, ok := <-result:
		if !ok {
			return "canceled"
		}
		return value
	case <-timer.C:
		return "timed out"
	}
}

func main() {}
