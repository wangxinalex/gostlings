// Concept: select timeout — stop waiting when no result becomes ready
// Task: return a ready result, or return "timed out" when the channel stays silent
// Expected behavior: a buffered result wins immediately; a silent channel times out
// Hint: use select with case value := <-ch and case <-time.After(...). The timeout
//       case must return "timed out" instead of waiting on ch forever.

package main

import "fmt"

func await(ch <-chan string) string {
	return "" // TODO: select between receiving ch and a timeout
}

func main() {
	result := make(chan string, 1)
	result <- "ready"
	fmt.Println(await(result))
}
