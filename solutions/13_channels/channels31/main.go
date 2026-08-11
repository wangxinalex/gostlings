package main

import "fmt"

func firstResult(tasks []func(<-chan struct{}) string) string {
	if len(tasks) == 0 {
		return ""
	}
	stop := make(chan struct{})
	winner := make(chan string, 1)
	exited := make(chan struct{}, len(tasks))
	for _, task := range tasks {
		go func(task func(<-chan struct{}) string) {
			defer func() { exited <- struct{}{} }()
			value := task(stop)
			select {
			case winner <- value:
			case <-stop:
			}
		}(task)
	}
	value := <-winner
	close(stop)
	for range tasks {
		<-exited
	}
	return value
}

func main() { fmt.Println(firstResult(nil)) }
