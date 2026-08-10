package main

import (
	"errors"
	"fmt"
	"sync"
)

type job struct {
	value int
	fail  bool
}

func run(workers int, jobs []job) error {
	if workers < 1 {
		workers = 1
	}

	jobsCh := make(chan job)
	stop := make(chan struct{})
	errorsCh := make(chan error, 1)
	var closeOnce sync.Once
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Go(func() {
			for {
				select {
				case <-stop:
					return
				case current, ok := <-jobsCh:
					if !ok {
						return
					}
					if current.fail {
						closeOnce.Do(func() {
							errorsCh <- errors.New("job failed")
							close(stop)
						})
						return
					}
				}
			}
		})
	}

	go func() {
		defer close(jobsCh)
		for _, current := range jobs {
			select {
			case jobsCh <- current:
			case <-stop:
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(errorsCh)
	}()

	var first error
	for err := range errorsCh {
		if first == nil {
			first = err
		}
	}
	return first
}

func main() {
	if err := run(2, []job{{value: 1}, {fail: true}, {value: 2}}); err != nil {
		fmt.Println(err)
	}
}
