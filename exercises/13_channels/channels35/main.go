// Concept: a cancellable worker group must protect both its receive and send operations.
// Task: square jobs with workers workers until jobs or stop closes.
// Expected behavior: cancellation releases blocked workers; a coordinator closes output after every worker exits.
// Hint: select on stop while receiving jobs and again while sending results. Workers acknowledge exit;
//       the coordinator receives every acknowledgement before close(out).
package main

import "fmt"

var onFanOutBeforeSend = func() {}

func fanOut(stop <-chan struct{}, jobs <-chan int, workers int) <-chan int {
	return nil // TODO: make receives and sends cancellable
}
func main() { fmt.Println(fanOut(make(chan struct{}), nil, 1)) }
