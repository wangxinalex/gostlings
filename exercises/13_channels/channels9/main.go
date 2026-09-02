// Problem: a caller wants to check for work now without waiting.
//
// Without this pattern: a normal receive blocks until a sender arrives.
//
// Channels: ch is a receive-only data input. default is the immediate fallback,
// not a value or completion signal.
//
// Timeline:
//   try receive
//   value ready -> receive
//   no value ready -> default -> return
//
// Hint:
//   Use select with a receive case and default. This is one immediate attempt;
//   do not turn default into a busy loop without useful work or backoff.

package main

import "fmt"

func tryReceive(ch <-chan int) string {
	// Thought: default makes select return immediately. It means “try now,” not
	// “guarantee that a value will eventually arrive.”
	return "" // TODO: add receive and default cases
}

func main() {
	fmt.Println(tryReceive(make(chan int)))
}
