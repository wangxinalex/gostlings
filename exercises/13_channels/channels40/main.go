// Concept: relay cancellation must cover a blocked input receive and a blocked output send.
// Task: forward values from in until in or stop closes.
// Expected behavior: normal values preserve order; stop closes output from either blocked direction.
// Hint: defer close(out). Select between stop and receiving in, check comma-ok, then select between
//       stop and sending the received value to out.
package main

import "fmt"

func relay(stop <-chan struct{}, in <-chan int) <-chan int {
	return nil // TODO: make both relay directions cancellation-aware
}
func main() { fmt.Println(relay(make(chan struct{}), nil)) }
