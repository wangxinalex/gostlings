// Concept: fan-in combines any number of input streams.
// Task: forward every input to one output and close that output only after every forwarder ends.
// Expected behavior: merge() forwards all values; merge() with no inputs returns a closed stream.
// Hint: start one goroutine per input and have each send one acknowledgement to a buffered exited channel.
//       The coordinator receives len(inputs) acknowledgements before it calls close(out).

package main

import "fmt"

func merge(inputs ...<-chan int) <-chan int {
	return nil // TODO: forward every input and close out from one coordinator
}

func main() {
	first := make(chan int, 1)
	second := make(chan int, 1)
	first <- 1
	second <- 2
	close(first)
	close(second)
	for value := range merge(first, second) {
		fmt.Println(value)
	}
}
