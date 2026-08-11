// Concept: a fan-in must handle stream lifecycle edges as well as normal values.
// Task: merge all supplied non-nil inputs, including already-closed and buffered streams.
// Expected behavior: no inputs closes immediately; buffered values drain before output closes.
// Hint: each forwarder ranges its own input so an already-closed input sends its acknowledgement immediately.
//       Size exited to len(inputs), and make one coordinator receive every acknowledgement before close(out).

package main

import "fmt"

var onForwarderExit = func() {}

func merge(inputs ...<-chan int) <-chan int {
	return nil // TODO: handle empty, closed, buffered, and still-open input streams
}

func main() {
	in := make(chan int, 2)
	in <- 1
	in <- 2
	close(in)
	for value := range merge(in) {
		fmt.Println(value)
	}
}
