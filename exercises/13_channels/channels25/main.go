// Problem: real fan-in inputs may be empty, already closed, buffered, or still
// producing; these states must not change the close protocol.
// Without this pattern: waiting for the wrong number of senders or forgetting a
// closed input can leave the output open forever.
// Channels: non-nil inputs carry data; out is coordinator-owned; exited has one
// acknowledgement slot per input. A closed input still needs its acknowledgement.
// Timeline: input drains (including buffered values) -> forwarder exits -> coordinator counts all -> close(out)
// Hint: handle no inputs explicitly. Start one forwarder per supplied input and
// range it so buffered values drain and already-closed inputs exit immediately.
// Size exited to len(inputs); the coordinator receives every acknowledgement
// before closing out. Do not let a forwarder close the shared output.

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
