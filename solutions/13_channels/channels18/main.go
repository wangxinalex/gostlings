package main

import "fmt"

func startWorkers(count int, stop <-chan struct{}) <-chan struct{} {
	done := make(chan struct{})
	if count <= 0 {
		close(done)
		return done
	}

	exited := make(chan struct{}, count)
	for worker := 0; worker < count; worker++ {
		go func() {
			<-stop
			exited <- struct{}{}
		}()
	}

	go func() {
		for worker := 0; worker < count; worker++ {
			<-exited
		}
		close(done)
	}()
	return done
}

func main() {
	stop := make(chan struct{})
	done := startWorkers(3, stop)
	close(stop)
	<-done
	fmt.Println("workers stopped")
}
