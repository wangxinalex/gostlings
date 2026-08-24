// Problem: callers should consume a stream without knowing how its producer is
// implemented or when it will finish.
// Without this pattern: the caller must coordinate the producer and may wait
// forever because nobody closes the output.
// Channels: out carries values; the producer owns sending and closing; callers
// receive and range over it.
// Timeline: create out -> producer sends values -> producer closes out -> caller's range ends
// Hint: create out, start the producer goroutine, defer close(out) inside that
// goroutine, and return out immediately. Empty input must close the stream too.

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
