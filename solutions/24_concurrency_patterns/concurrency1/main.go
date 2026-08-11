// Concept: a coordinator owns channel closure in a small pipeline.
// Task: generate every value, close the output, and sum until input closes.
// Hint: the producer that creates the channel should close it; sum should range over it.
package main

func generate(values ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range values {
			out <- value
		}
	}()
	return out
}

func sum(in <-chan int) int {
	total := 0
	for value := range in {
		total += value
	}
	return total
}

func main() {}
