// Concept: a consumer that abandons a producer's output must be able to release the blocked send.
// Task: produce a short sequence until it finishes or stop closes.
// Expected behavior: values are produced in order; closing stop lets the producer close its output even when
// nobody is receiving it anymore.
// Hint: create out and defer close(out) in one goroutine. For every value, select between out <- value and
// <-stop; never leave a producer blocked on an abandoned output send.
package main

import "fmt"

func produce(stop <-chan struct{}) <-chan int {
	return nil // TODO: make each output send stop-aware and close out on every exit path
}

func main() { fmt.Println(produce(make(chan struct{}))) }
