// Concept: generator — the producer owns and closes its output channel
// Task: send every input value from a goroutine, then close the returned channel
// Expected behavior: callers can range over the result, including for empty input
// Hint: defer close(out) inside the producer goroutine; the caller only receives

package main

import "fmt"

func generate(values ...int) <-chan int {
	// Thought: return a receive-only channel and keep sending and closing inside
	// the producer; callers only range over it and cannot close it accidentally.
	return nil // TODO: create the output, send values in a goroutine, and close it
}

func main() {
	for value := range generate(1, 2, 3) {
		fmt.Println(value)
	}
}
