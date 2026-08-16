// Concept: a service can select between control commands, work, and shutdown.
// Task: forward jobs unless pause is active; resume accepts jobs again; stop closes output.
// Expected behavior: paused jobs wait for resume, and stop wins even while jobs are blocked.
// Hint: model the service as two states. Initially it is active:
//       select among stop, commands, and jobs; apply commands and forward jobs.
//       When pause is received, set paused=true and stop receiving jobs. While paused,
//       select only stop and commands; resume sets paused=false. Do not use a default case,
//       or the loop will busy-spin and may stop observing commands fairly.
//       When forwarding a job, use a second select between out <- job and <-stop because
//       the downstream send is another blocking point. A closed commands or jobs channel
//       ends the service; this goroutine is the only owner that closes out.
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
