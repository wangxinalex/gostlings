// Concept: fan-in — merge multiple input channels into one output
// Task: forward all values and close the merged output after every input finishes
// Expected behavior: ranging over merge(...) terminates for normal and empty input
// Hint: one forwarder per input, one WaitGroup, and one goroutine that closes out

package main

import "fmt"

func merge(inputs ...<-chan int) <-chan int {
	// Thought: forwarders only send and must not close out independently. A
	// coordinator waits for all forwarders, then closes out once.
	return nil // TODO: start forwarders, wait for all, and close out once
}

func main() {
	left := make(chan int, 2)
	right := make(chan int, 2)
	left <- 1
	left <- 3
	right <- 2
	right <- 4
	close(left)
	close(right)
	for value := range merge(left, right) {
		fmt.Println(value)
	}
}
