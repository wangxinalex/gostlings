// Concept: select multiplexes channel operations
// Task: receive from whichever input is ready without blocking on a silent input
// Expected behavior: a ready input is returned; if multiple inputs are ready, either may be selected
// Hint: put one receive from each input in a select. select does not give earlier
//       cases priority when multiple cases are ready.

package main

import "fmt"

func receiveFast(fast, slow <-chan string) string {
	// Thought: select waits for multiple channel operations at once; do not
	// receive from slow first because it may never have a sender.
	// If both channels are ready, select chooses one without guaranteeing
	// which one wins.
	return <-fast // TODO: wait for whichever input is ready first
}

func main() {
	fast := make(chan string, 1)
	slow := make(chan string)
	fast <- "fast lane"
	fmt.Println(receiveFast(fast, slow))
}
