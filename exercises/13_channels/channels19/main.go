// Concept: timeout, cancellation, and join — returning is safe only after cleanup
// Task: time out a slow producer, close stop, wait for done, and return "timed out"
// Expected behavior: run returns after the producer has observed stop and closed done
// Hint: keep result data and completion notification separate. The intended sequence is:
//       stop := make(chan struct{}); result := make(chan string, 1)
//       go func() { defer close(done); runProducer(stop, result) }()
//       select on result or time.After(25 * time.Millisecond)
//       in either branch: close(stop), then receive <-done before returning.
//       The producer selects between its slow work and <-stop. close(stop) is the request
//       to stop; <-done is the join that confirms the producer has actually stopped.

package main

import "fmt"

var runProducer = func(stop <-chan struct{}, result chan<- string) {}

func run(done chan struct{}) string {
	return "" // TODO: cancel the producer on timeout and join it before returning
}

func main() {
	done := make(chan struct{})
	fmt.Println(run(done))
}
