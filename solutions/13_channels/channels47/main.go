package main

import "fmt"

type response struct {
	value int
	err   error
}

type request struct {
	value int
	reply chan response
	err   error
}

const serveWorkers = 2

var onServeBeforeResult = func() {}

func serve(stop <-chan struct{}, jobs <-chan request) (<-chan response, <-chan struct{}) {
	results := make(chan response)
	done := make(chan struct{})
	exited := make(chan struct{}, serveWorkers)
	for range serveWorkers {
		go func() {
			defer func() { exited <- struct{}{} }()
			for {
				select {
				case <-stop:
					return
				default:
				}
				var current request
				select {
				case <-stop:
					return
				case next, ok := <-jobs:
					if !ok {
						return
					}
					current = next
				}
				select {
				case <-stop:
					return
				default:
				}
				answer := response{value: current.value * 2, err: current.err}
				if current.reply != nil {
					select {
					case <-stop:
						return
					case current.reply <- answer:
					}
				}
				onServeBeforeResult()
				select {
				case <-stop:
					return
				case results <- answer:
				}
			}
		}()
	}
	go func() {
		for range serveWorkers {
			<-exited
		}
		close(results)
		close(done)
	}()
	return results, done
}

func main() { fmt.Println(serve(make(chan struct{}), nil)) }
