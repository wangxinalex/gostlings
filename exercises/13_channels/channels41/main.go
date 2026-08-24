// Advanced pattern: nil-channel switching disables a select case dynamically.
// Problem: a closed channel is always ready, even after its buffered values are
// drained.
// Without this pattern: select repeatedly chooses the closed input and returns
// its zero value, starving the other input or spinning forever.
// Channels: first and second carry data; assigning a local input to nil removes
// that case from select. No function here owns either input.
// Timeline: receive until ok=false -> set that input to nil -> drain other input -> return
// Hint: loop while first != nil || second != nil. Use comma-ok in each receive
// case; when !ok, set only that local input variable to nil. Append only values
// received with ok=true.
package main

import "fmt"

func drain(first, second <-chan int) []int {
	return nil // TODO: disable each closed input with nil after comma-ok reports closure
}

func main() { fmt.Println(drain(nil, nil)) }
