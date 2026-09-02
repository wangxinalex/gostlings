// Problem: one input may be ready while another input is silent forever.
//
// Without this pattern: receiving from the silent input first can block even
// though useful data is already available elsewhere.
//
// Channels: fast and slow are receive-only data inputs. No channel is closed by
// this function.
//
// Timeline:
//   select waits on fast and slow
//   first ready receive wins
//
// Hint:
//   Put one receive case per input in select. If several cases are ready,
//   select chooses among them without giving source order priority.

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
