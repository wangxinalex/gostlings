// Concept: channel ownership and directional channel types
// Task: make the producer send every value and close its output; make sum drain until close
// Expected output: focused channel tests pass (run `go test ./solutions/24_concurrency_patterns/pipeline1`)
// Hint: the goroutine that creates a channel should close it; range over <-chan int in sum

package main

func generate(values ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range values {
			out <- value
		}
	}()
	return out
}

func sum(in <-chan int) int {
	total := 0
	for value := range in {
		total += value
	}
	return total
}
