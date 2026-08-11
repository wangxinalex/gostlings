package main

import "fmt"

var onFanOutBeforeSend = func() {}

func fanOut(stop <-chan struct{}, jobs <-chan int, workers int) <-chan int {
	out := make(chan int)
	if workers < 1 {
		close(out)
		return out
	}
	exited := make(chan struct{}, workers)
	for range workers {
		go func() {
			defer func() { exited <- struct{}{} }()
			for {
				var job int
				select {
				case <-stop:
					return
				case value, ok := <-jobs:
					if !ok {
						return
					}
					job = value
				}
				onFanOutBeforeSend()
				select {
				case <-stop:
					return
				default:
				}
				select {
				case <-stop:
					return
				case out <- job * job:
				}
			}
		}()
	}
	go func() {
		for range workers {
			<-exited
		}
		close(out)
	}()
	return out
}

func main() { fmt.Println(fanOut(make(chan struct{}), nil, 1)) }
