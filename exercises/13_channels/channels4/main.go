// Concept: comma-ok receive distinguishes a closed channel from a real zero value
// Task: use the receive form that reports whether a value was actually received
// Expected behavior: a closed channel reports value 0 with ok=false
// Hint: use value, ok := <-ch; ok is false only after the channel is closed and drained

package main

import "fmt"

func read(ch <-chan int) (int, bool) {
	// Thought: when receiving an int, zero may be real data or the value returned
	// after closure; the comma-ok result is what identifies the channel state.
	// Pattern:
	//   value, ok := <-ch
	//   if !ok { // channel is closed and drained
	//       stop consuming
	//   }
	value := <-ch
	return value, true // TODO: receive with comma-ok and return the real status
}

func main() {
	ch := make(chan int)
	close(ch)
	value, ok := read(ch)
	fmt.Println(value, ok)
}
