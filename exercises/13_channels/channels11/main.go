// Problem: a caller cannot wait forever for a result from an unhealthy or slow
// operation.
//
// Without this pattern: a silent result channel blocks the caller indefinitely.
//
// Channels: ch carries the result. time.After(100*time.Millisecond) provides a
// deadline signal. This function owns neither channel.
//
// Timeline:
//   wait for result and deadline
//   result ready -> return early
//   deadline ready at 100ms -> return timeout
//
// Hint:
//   Wait for the result and a deadline at the same time: add a second select
//   case whose channel becomes ready after the configured duration. This is a
//   maximum wait, not a mandatory 100ms sleep.
// Stuck?: time.After produces a one-shot channel; use exactly 100 milliseconds.

package main

import "fmt"

func await(ch <-chan string) string {
	// The maximum wait is 100ms. If ch receives a value first, return immediately;
	// this is not a mandatory 100ms sleep. Only a silent ch reaches the timeout case.
	// time.After returns a channel that becomes ready at the deadline, so it can
	// participate in the same select as the result channel.
	return "" // TODO: select between receiving ch and a timeout
}

func main() {
	result := make(chan string, 1)
	result <- "ready"
	fmt.Println(await(result))
}
