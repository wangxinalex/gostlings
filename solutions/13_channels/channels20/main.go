package main

import "fmt"

func shutdown(stop chan struct{}, workers int) <-chan struct{} {
	done := make(chan struct{})
	if workers < 0 {
		workers = 0
	}

	exited := make(chan struct{}, workers)
	for worker := 0; worker < workers; worker++ {
		go func() {
			<-stop
			exited <- struct{}{}
		}()
	}

	go func() {
		select {
		case <-stop:
		default:
			close(stop)
		}
		for worker := 0; worker < workers; worker++ {
			<-exited
		}
		close(done)
	}()
	return done
}

func main() {
	stop := make(chan struct{})
	<-shutdown(stop, 3)
	fmt.Println("shutdown complete")
}
