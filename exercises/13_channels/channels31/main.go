// Concept: the first completed task can cancel its peers.
// Task: return the first task result, broadcast stop, and wait for every task to exit.
// Expected behavior: a task that completes first wins; tasks still running observe stop before return.
// Hint: give every task the same stop channel. Send results through a capacity-one channel, then
//
//	close stop once and receive one exit acknowledgement for each task before returning.
package main

import "fmt"

func firstResult(tasks []func(<-chan struct{}) string) string {
	return "" // TODO: publish one winner, broadcast stop, and join every task
}

func main() { fmt.Println(firstResult(nil)) }
