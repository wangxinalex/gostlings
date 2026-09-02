// Problem: a semaphore limits concurrency, but cancellation can arrive while a
// worker waits for capacity or while it publishes its result.
// Without this pattern: one worker can remain blocked on tokens or results after
// the caller has stopped waiting.
// Channels: tokens carries permits; stop requests cancellation; results carries
// indexed data; exited/done join every started worker before return.
// Timeline:
//   select on token or stop
//   perform work
//   release the token
//   select on result or stop
//   join workers
//   return
//
// Hint:
//   Select between `<-stop` and `<-tokens`.
//   Skip the job if stop wins.
//   Run work and return the token before publishing.
//   Select between stop and the result send.
//   A coordinator waits for every started goroutine.
//   The coordinator closes results before the collector finishes.
//   Return false only after joining cancellation.

package main

import "fmt"

func parallel(stop <-chan struct{}, limit int, jobs []int, work func(int) int) ([]int, bool) {
	return nil, false // TODO: make token acquisition and result publication cancellable
}

func main() { fmt.Println(parallel(make(chan struct{}), 1, nil, func(value int) int { return value })) }
