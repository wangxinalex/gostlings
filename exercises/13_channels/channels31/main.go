// Concept: the first completed task can cancel its peers.
// Task: return the first task result, broadcast stop, and wait for every task to exit.
// Expected behavior: a task that completes first wins; tasks still running observe stop before return.
// Hint: give every task the same stop channel and one exit acknowledgement slot. Each task computes
//       its result, publishes it through a capacity-one result channel, and acknowledges exit even
//       when it observes stop. The caller waits for the first result, closes stop exactly once, then
//       receives one exit acknowledgement per task before returning the winner. A task must honor
//       stop itself; closing a channel cannot interrupt arbitrary CPU work or a non-channel call.
//       With no tasks, return the empty string without starting any goroutine.
package main

import "fmt"

func firstResult(tasks []func(<-chan struct{}) string) string {
	return "" // TODO: publish one winner, broadcast stop, and join every task
}

func main() { fmt.Println(firstResult(nil)) }
