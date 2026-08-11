// Concept: a service can select between control commands, work, and shutdown.
// Task: forward jobs unless pause is active; resume accepts jobs again; stop closes output.
// Expected behavior: paused jobs wait for resume, and stop wins even while jobs are blocked.
// Hint: while paused, select only stop and commands. While active, select stop, commands, and jobs;
//
//	when forwarding a job, select between sending it and stop. This service owns close(out).
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
