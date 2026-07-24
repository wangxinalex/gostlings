// Concept: custom type constraints and generic functions over slices
// Task: define a Number constraint and a generic Sum function so this program compiles
// Expected output: 4
// Hint: use ~int | ~float64 in the interface to allow both types (Go Tour: Generics 1)

package main

import "fmt"

// TODO: Define a Number constraint (interface with ~int | ~float64).

// TODO: Define Sum[T Number](xs []T) T that adds up all elements.

func main() {
	fmt.Println(Sum([]float64{1.5, 2.5}))
}
