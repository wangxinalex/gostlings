// Concept: a closed channel stays ready forever, so a select must disable it after comma-ok says it is closed.
// Task: drain both inputs until both close.
// Expected behavior: buffered values arrive once; after closure, set that input variable to nil so it cannot
// repeatedly supply its zero value.
// Hint: loop while first != nil || second != nil. In each receive case, use value, ok := <-input; when !ok,
// set only that local input variable to nil and continue. Append only values received with ok == true.
package main

import "fmt"

func drain(first, second <-chan int) []int {
	return nil // TODO: disable each closed input with nil after comma-ok reports closure
}

func main() { fmt.Println(drain(nil, nil)) }
