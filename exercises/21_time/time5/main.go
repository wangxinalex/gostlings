// Concept: select can distinguish a result, a closed result stream, and a timeout.
// Task: return result, canceled, or timed out without leaking the timer.
// Hint: use a timer and check the receive's ok value for the cancellation branch.
package main

import "time"

func awaitOrCancel(result <-chan string, timeout time.Duration) string {
	// TODO: Distinguish a value, a closed result channel, and timer expiration.
	return ""
}

func main() {}
