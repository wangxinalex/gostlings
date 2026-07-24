// Concept: custom type constraints and generic functions over slices
// Task: define a Number constraint and a generic Sum function so this program compiles
// Expected output: 4
// Hint: use ~int | ~float64 in the interface to allow both types (Go Tour: Generics 1)

package main

import "fmt"

type Number interface {
	~int | ~float64
}

func Sum[T Number](xs []T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total
}

func main() {
	fmt.Println(Sum([]float64{1.5, 2.5}))
}
