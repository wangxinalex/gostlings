package main

import "fmt"

type job struct {
	value int
	err   error
}

var onStopClosed = func() {}
var onWorkerExit = func() {}

func run(workers int, jobs []job) error {
	if workers < 1 {
		workers = 1
	}

	stop := make(chan struct{})
	jobsCh := make(chan job)
	failures := make(chan error, workers)
	exited := make(chan struct{}, workers)
	go func() {
		defer close(jobsCh)
		for _, job := range jobs {
			select {
			case jobsCh <- job:
			case <-stop:
				return
			}
		}
	}()
	for worker := 0; worker < workers; worker++ {
		go func() {
			var failure error
			for {
				select {
				case <-stop:
					onWorkerExit()
					exited <- struct{}{}
					return
				case current, ok := <-jobsCh:
					if !ok {
						onWorkerExit()
						exited <- struct{}{}
						return
					}
					if current.err != nil {
						failure = current.err
						failures <- failure
						onWorkerExit()
						exited <- struct{}{}
						return
					}
				}
			}
		}()
	}

	var first error
	stopClosed := false
	closeStop := func() {
		if stopClosed {
			return
		}
		close(stop)
		stopClosed = true
		onStopClosed()
	}
	for remaining := workers; remaining > 0; {
		select {
		case failure := <-failures:
			if first == nil {
				first = failure
				closeStop()
			}
		case <-exited:
			remaining--
		}
	}
	for {
		select {
		case failure := <-failures:
			if first == nil {
				first = failure
				closeStop()
			}
		default:
			return first
		}
	}
}

func main() {
	fmt.Println(run(2, []job{{value: 1}, {value: 2}}))
}
