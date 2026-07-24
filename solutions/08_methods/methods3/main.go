// Concept: defining methods with a pointer receiver
// Task: define the Scale method so this program compiles and runs
// Expected output: 24
// Hint: a pointer receiver (*Rectangle) lets the method modify the original value (Go Tour: Methods 4)

package main

import "fmt"

type Rectangle struct {
	W, H int
}

func (r Rectangle) Area() int {
	return r.W * r.H
}

func (r *Rectangle) Scale(f int) {
	r.W *= f
	r.H *= f
}

func main() {
	r := Rectangle{2, 3}
	r.Scale(2)
	fmt.Println(r.Area())
}
