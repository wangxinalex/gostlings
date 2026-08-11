package main

import "fmt"

type command int

const (
	pause command = iota
	resume
)

var onCommandApplied = func(command) {}

func serveCommands(stop <-chan struct{}, commands <-chan command, jobs <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		paused := false
		for {
			if paused {
				select {
				case <-stop:
					return
				case current, ok := <-commands:
					if !ok {
						return
					}
					if current == resume {
						paused = false
					}
					onCommandApplied(current)
				}
				continue
			}
			select {
			case <-stop:
				return
			case current, ok := <-commands:
				if !ok {
					return
				}
				if current == pause {
					paused = true
				}
				onCommandApplied(current)
			case job, ok := <-jobs:
				if !ok {
					return
				}
				select {
				case <-stop:
					return
				case out <- job:
				}
			}
		}
	}()
	return out
}
func main() { fmt.Println(serveCommands(make(chan struct{}), nil, nil)) }
