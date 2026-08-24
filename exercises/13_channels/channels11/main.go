// Problem: a caller cannot wait forever for a result from an unhealthy or slow
// operation.
// Without this pattern: a silent result channel blocks the caller indefinitely.
// Channels: ch carries the result; time.After(100*time.Millisecond) provides a
// deadline signal; this function owns neither channel.
// Timeline: wait for result and deadline -> result wins early, or deadline wins at 100ms
// Hint: use exactly 100 milliseconds in a select with the result receive and
// `<-time.After(...)`. This is a maximum wait, not a mandatory 100ms sleep.

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
