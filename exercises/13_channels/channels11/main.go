// Concept: select timeout — stop waiting when no result becomes ready
// Task: return a ready result, or return "timed out" when the channel stays silent
// Expected behavior: a buffered result wins immediately; a silent channel times out
// Hint: wait at most 100 milliseconds: use select with case value := <-ch and
//       case <-time.After(100 * time.Millisecond). The timeout case must return
//       "timed out" instead of waiting on ch forever.

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
