// Concept: a consumer that abandons an output must signal upstream work to stop.
// Task: forward work until work closes or stop closes.
// Expected behavior: normal values arrive in order; stop releases a forwarder blocked sending downstream.
// Hint: defer close(out). Use one select for the work receive and another for the out send, with a
//       stop case in both; only this forwarding goroutine closes out.
package main

import "fmt"

func collectOrStop(stop <-chan struct{}, work <-chan int) <-chan int {
	return nil // TODO: make the receive and send cancellable
}
func main() { fmt.Println(collectOrStop(make(chan struct{}), nil)) }
