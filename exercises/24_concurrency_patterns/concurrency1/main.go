// Concept: a coordinator owns channel closure in a small pipeline.
// Task: generate every value, close the output, and sum until input closes.
// Hint: the producer that creates the channel should close it; sum should range over it.
package main

func generate(values ...int) <-chan int {
	out := make(chan int)
	// TODO: Send every value from a goroutine and close out exactly once.
	return out
}

func sum(in <-chan int) int {
	// TODO: Drain in until close and return the total.
	return 0
}
