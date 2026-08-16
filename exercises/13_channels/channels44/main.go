// Concept: cancellation must cover both waiting for semaphore capacity and publishing worker results.
// Task: run bounded work until completion or stop closes.
// Expected behavior: false reports cancellation; no worker remains blocked acquiring a token or sending a result.
// Hint: use one semaphore token per active call. For each indexed job:
//
//	select between <-stop and <-tokens; if stop wins, skip the job.
//	Run work(value), then return the token before publishing the result.
//	Select between <-stop and results <- indexedResult so a blocked collector can be canceled.
//	A coordinator waits for every started goroutine, closes results, and only then lets the
//	collector finish. Store results by index; return false if stop was observed, after joining.
package main

import "fmt"

func parallel(stop <-chan struct{}, limit int, jobs []int, work func(int) int) ([]int, bool) {
	return nil, false // TODO: make token acquisition and result publication cancellable
}

func main() { fmt.Println(parallel(make(chan struct{}), 1, nil, func(value int) int { return value })) }
