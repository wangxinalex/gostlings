// Concept: fan-in — merge multiple input channels into one output channel
// Task: forward every input value and close the merged output after all inputs finish
// Expected behavior: all values arrive, then ranging over the result terminates
// Hint: one forwarding goroutine per input, sync.WaitGroup, and one closer goroutine

package main

func merge(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	// TODO: Start forwarders, wait for all inputs, and close out exactly once.
	return out
}
