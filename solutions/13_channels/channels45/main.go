package main

import "fmt"

func rateLimit(tokens <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			<-tokens
			out <- value
		}
	}()
	return out
}

func main() { fmt.Println(rateLimit(nil, nil)) }
