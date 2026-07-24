// Concept: generic functions with type parameters and constraints
// Task: define a generic Max function so this program compiles and runs
// Expected output: 7
// 2.5
// Hint: use [T cmp.Ordered] as the type parameter list and compare with > (Go Tour: Generics 1)

package main

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(3, 7))
	fmt.Println(Max(2.5, 1.5))
}
