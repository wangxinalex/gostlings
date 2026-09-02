// Problem: several equivalent tasks race, but the caller only needs the first
// useful answer and must stop the losing tasks.
// Without this pattern: losers continue consuming resources, and returning the
// winner without joining them leaks goroutines.
// Channels: result carries task answers; stop broadcasts cancellation; exited
// carries one acknowledgement per task. Each task must observe stop at its own
// blocking points; closing a channel cannot interrupt arbitrary CPU work.
// Timeline:
//   tasks race
//   first result arrives
//   close(stop)
//   every task sends one exit acknowledgement
//   return the winner
//
// Hint:
//   Give every task the same stop channel and one exit acknowledgement slot.
//   Publish results through a capacity-one channel.
//   Acknowledge exit on every path.
//   Close stop once after receiving the first result.
//   Receive one exit token per task before returning.
//   With no tasks, return the empty string immediately.

package main

import "fmt"

func firstResult(tasks []func(<-chan struct{}) string) string {
	return "" // TODO: publish one winner, broadcast stop, and join every task
}

func main() { fmt.Println(firstResult(nil)) }
