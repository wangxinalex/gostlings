// Concept: nil-channel switching retires an input that has closed.
// Task: drain both inputs into one slice without losing buffered values.
// Expected behavior: every buffered value is returned exactly once; an input that
// is already closed contributes nothing and does not keep the loop alive.
// Hint: a closed channel is always ready, so a loop that leaves it selectable
//       receives its zero value forever instead of moving on; a nil channel is
//       never ready, so assigning nil removes that case from the select. Comma-ok
//       reports the closure.
// Stuck?: keep the two input variables in the loop condition and clear only the
//       one that reported a closed receive.

package main

import "fmt"

func drain(first, second <-chan int) []int {
	// TODO: collect every buffered value from both inputs, disabling each input
	//       once a receive reports that it is closed.
	return nil
}

func main() { fmt.Println(drain(nil, nil)) }
