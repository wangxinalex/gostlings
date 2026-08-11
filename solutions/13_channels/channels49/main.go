package main

import "fmt"

type job struct {
	value int
	err   error
}

var processFirstErrorBounded = func(value int) int { return value * value }

func runFirstErrorBounded(stop <-chan struct{}, workers int, jobs []job) ([]int, error) {
	if workers < 1 {
		workers = 1
	}
	cancel := make(chan struct{})
	jobsCh := make(chan job)
	results := make(chan int)
	failures := make(chan error, 1)
	exited := make(chan struct{}, workers)
	go func() {
		defer close(jobsCh)
		for _, current := range jobs {
			select {
			case <-stop:
				return
			case <-cancel:
				return
			default:
			}
			select {
			case <-stop:
				return
			case <-cancel:
				return
			case jobsCh <- current:
			}
		}
	}()
	for range workers {
		go func() {
			defer func() { exited <- struct{}{} }()
			for {
				select {
				case <-stop:
					return
				case <-cancel:
					return
				default:
				}
				var current job
				select {
				case <-stop:
					return
				case <-cancel:
					return
				case next, ok := <-jobsCh:
					if !ok {
						return
					}
					current = next
				}
				if current.err != nil {
					select {
					case <-stop:
					case <-cancel:
					case failures <- current.err:
					}
					return
				}
				value := processFirstErrorBounded(current.value)
				select {
				case <-stop:
					return
				case <-cancel:
					return
				case results <- value:
				}
			}
		}()
	}
	go func() {
		for range workers {
			<-exited
		}
		close(results)
	}()

	values := []int{}
	var first error
	canceled := false
	stopCh := stop
	closeCancel := func() {
		if !canceled {
			close(cancel)
			canceled = true
		}
	}
	for results != nil {
		select {
		case value, ok := <-results:
			if !ok {
				results = nil
				continue
			}
			values = append(values, value)
		case err := <-failures:
			if first == nil {
				first = err
				closeCancel()
			}
		case <-stopCh:
			closeCancel()
			stopCh = nil
		}
	}
	for {
		select {
		case err := <-failures:
			if first == nil {
				first = err
				closeCancel()
			}
		default:
			return values, first
		}
	}
}

func main() { fmt.Println(runFirstErrorBounded(make(chan struct{}), 1, nil)) }
