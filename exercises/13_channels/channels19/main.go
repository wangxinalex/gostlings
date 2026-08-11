// Concept: timeout, cancellation, and join — returning is safe only after cleanup
// Task: time out a slow producer, close stop, wait for done, and return "timed out"
// Expected behavior: run returns after the producer has observed stop and closed done
// Hint: the producer should defer close(done) and select on its slow delay or stop. The
//       caller selects on result or time.After; in either branch close stop, then receive done.

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
