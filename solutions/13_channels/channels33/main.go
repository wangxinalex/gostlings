package main

import "fmt"

var onCollectorExit = func() {}

func collect(sources []<-chan int) <-chan int {
	out := make(chan int)
	active := 0
	for _, source := range sources {
		if source != nil {
			active++
		}
	}
	if active == 0 {
		close(out)
		return out
	}
	exited := make(chan struct{}, active)
	for _, source := range sources {
		if source == nil {
			continue
		}
		go func(source <-chan int) {
			for value := range source {
				out <- value
			}
			onCollectorExit()
			exited <- struct{}{}
		}(source)
	}
	go func() {
		for range active {
			<-exited
		}
		close(out)
	}()
	return out
}

func main() { fmt.Println(collect(nil)) }
