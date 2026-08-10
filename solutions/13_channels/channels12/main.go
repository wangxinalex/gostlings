package main

import (
	"fmt"
	"sync"
)

func startWorkers(count int, stop <-chan struct{}) <-chan struct{} {
	done := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Go(func() {
			<-stop
		})
	}

	go func() {
		wg.Wait()
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
