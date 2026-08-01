// Concept: channel ownership and directional channel types
// Task: make the producer send every value and close its output; make sum drain until close
// Expected output: focused channel tests pass (run `go test ./exercises/24_concurrency_patterns/pipeline1`)
// Hint: the goroutine that creates a channel should close it; range over <-chan int in sum

package main

func generate(values ...int) <-chan int {
	out := make(chan int)
	// TODO: Start a goroutine that sends values to out and closes out.
	return out
}

func sum(in <-chan int) int {
	// TODO: Range over in, accumulate all values, and return the total after close.
	return 0
}
