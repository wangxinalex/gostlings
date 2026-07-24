// Concept: pointer receivers mutate the original value
// Task: this program should print 1, but it prints 0; fix the Increment method
// Expected output: 1
// Hint: methods with value receivers operate on a copy; use a pointer receiver to modify the receiver (Go Tour: Methods 4)

package main

import "fmt"

type Counter struct {
	n int
}

func (c *Counter) Increment() {
	c.n++
}

func (c Counter) Value() int {
	return c.n
}

func main() {
	c := Counter{}
	c.Increment()
	fmt.Println(c.Value())
}
