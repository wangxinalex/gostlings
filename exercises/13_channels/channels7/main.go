// Problem: a generator should prevent callers from accidentally sending to or
// closing the producer's output.
//
// Without this pattern: a bidirectional return type exposes operations that do
// not belong to the caller and can cause a panic or protocol violation.
//
// Channels: the producer uses bidirectional out internally. The caller receives
// through the returned <-chan int. The producer owns close(out).
//
// Timeline:
//   producer sends -> caller receives
//   producer closes -> caller's range ends
//
// Hint:
//   Make a bidirectional channel inside generate, start the producer with
//   defer close(out), and return the same channel as <-chan int.

package main

import "fmt"

func receiveAll(ch <-chan int) []int {
	var values []int
	for value := range ch {
		values = append(values, value)
	}
	return values
}

func generate(values ...int) <-chan int {
	// Thought: callers receive through the returned <-chan int, so only this
	// producer can send values and decide when the stream is complete.
	return nil // TODO: create out, send values in a goroutine, close it, and return it
}

func main() {
	for _, value := range receiveAll(generate(1, 2, 3)) {
		fmt.Println(value)
	}
}
