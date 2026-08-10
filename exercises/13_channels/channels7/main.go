// Concept: select timeout — a caller should not wait forever for a result
// Task: return the value if it arrives quickly, otherwise return a timeout message
// Expected behavior: a silent input returns "timed out" promptly
// Hint: add a case for <-time.After(50 * time.Millisecond)

package main

import (
	"fmt"
	"time"
)

func await(ch <-chan string) string {
	// Thought: a timeout is one competing select branch. It ends only the current
	// wait; it does not automatically stop a producer that is still running.
	timeout := time.After(50 * time.Millisecond)
	_ = timeout
	return <-ch // TODO: add the timeout branch
}

func main() {
	fmt.Println(await(make(chan string)))
}
