// Problem: receiving an int from a closed channel also produces the zero value.
//
// Without this pattern: value == 0 cannot tell real data from a closed stream.
//
// Channels: ch is a receive-only data channel. ok reports whether a value was
// received before closure.
//
// Timeline:
//   close(ch) -> receive value=0, ok=false
//
// Hint:
//   Use the comma-ok receive form. ok is false only when ch is closed and
//   drained; do not infer channel state from the value alone.

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
