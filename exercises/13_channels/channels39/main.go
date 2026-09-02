// Problem: a long-lived service must handle control commands and work while
// remaining cancellable.
// Without this pattern: pause may still consume jobs, or a blocked downstream
// send can prevent stop from taking effect.
// Channels: commands carries control state changes; jobs carries data; stop is
// cancellation; out is service-owned and closed by this goroutine.
// Timeline:
//   active select: commands, jobs, or stop
//   pause disables jobs
//   resume restores jobs
//   stop ends the service
//   service closes out
//
// Hint:
//   Model the service as two states.
//   Initially it is active:
//
//       select among stop, commands, and jobs
//       apply commands and forward jobs
//
//   When pause is received, set paused=true and stop receiving jobs.
//   While paused, select only stop and commands.
//   Resume sets paused=false.
//   Do not use a default case, or the loop will busy-spin and may stop
//   observing commands fairly.
//
//   When forwarding a job, use a second select between out <- job and <-stop.
//   The downstream send is another blocking point.
//   A closed commands or jobs channel ends the service.
//   This goroutine is the only owner that closes out.

package main

import "fmt"

type command int

const (
	pause command = iota
	resume
)

var onCommandApplied = func(command) {}

func serveCommands(stop <-chan struct{}, commands <-chan command, jobs <-chan int) <-chan int {
	return nil // TODO: select control, data, and shutdown without starving stop
}
func main() { fmt.Println(serveCommands(make(chan struct{}), nil, nil)) }
