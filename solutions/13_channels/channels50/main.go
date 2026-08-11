package main

import (
	"errors"
	"fmt"
)

type response struct {
	value int
	err   error
}

type request struct {
	value int
	reply chan response
	err   error
}

type indexedRequest struct {
	index   int
	request request
}

type indexedResponse struct {
	index    int
	response response
}

var errStopped = errors.New("service stopped")
var processServiceRequest = func(current request) response {
	return response{value: current.value * 2, err: current.err}
}

func runService(stop <-chan struct{}, workers int, requests []request) ([]response, error) {
	if workers < 1 {
		workers = 1
	}
	select {
	case <-stop:
		return []response{}, errStopped
	default:
	}
	if len(requests) == 0 {
		return []response{}, nil
	}

	cancel := make(chan struct{})
	jobs := make(chan indexedRequest)
	results := make(chan indexedResponse)
	failures := make(chan error, 1)
	exited := make(chan struct{}, workers)
	go func() {
		defer close(jobs)
		for index, current := range requests {
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
			case jobs <- indexedRequest{index: index, request: current}:
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
				var current indexedRequest
				select {
				case <-stop:
					return
				case <-cancel:
					return
				case next, ok := <-jobs:
					if !ok {
						return
					}
					current = next
				}
				answer := processServiceRequest(current.request)
				if current.request.err != nil {
					answer.err = current.request.err
				}
				if current.request.reply != nil {
					select {
					case <-stop:
						return
					case <-cancel:
						return
					case current.request.reply <- answer:
					}
				}
				if answer.err != nil {
					select {
					case <-stop:
					case <-cancel:
					case failures <- answer.err:
					}
					return
				}
				select {
				case <-stop:
					return
				case <-cancel:
					return
				case results <- indexedResponse{index: current.index, response: answer}:
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

	ordered := make([]response, len(requests))
	seen := make([]bool, len(requests))
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
		case result, ok := <-results:
			if !ok {
				results = nil
				continue
			}
			ordered[result.index] = result.response
			seen[result.index] = true
		case err := <-failures:
			if first == nil {
				first = err
				closeCancel()
			}
		case <-stopCh:
			if first == nil {
				first = errStopped
			}
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
			responses := make([]response, 0, len(requests))
			for index, response := range ordered {
				if seen[index] {
					responses = append(responses, response)
				}
			}
			return responses, first
		}
	}
}

func main() { fmt.Println(runService(make(chan struct{}), 1, nil)) }
