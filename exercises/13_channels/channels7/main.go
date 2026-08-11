// Concept: directional channels keep generator ownership clear
// Task: return a receive-only stream while the producer sends every value and closes its own output
// Expected behavior: callers can receive every value but cannot send to or close the returned channel
// Hint: make a bidirectional out channel inside generate, start a goroutine with defer close(out), and return it as <-chan int

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
