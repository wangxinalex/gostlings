// Concept: collecting one error per job and reporting it after a full join
// Task: run every job concurrently, wait for all of them, and return the first non-nil error

package main

import "sync"

func runAll(jobs []func() error) error {
	errs := make([]error, len(jobs))
	var wg sync.WaitGroup
	for index, job := range jobs {
		wg.Add(1)
		go func(index int, job func() error) {
			defer wg.Done()
			errs[index] = job()
		}(index, job)
	}
	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {}
