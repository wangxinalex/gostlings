package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	for value := range ch {
		fmt.Println(value)
	}
}
