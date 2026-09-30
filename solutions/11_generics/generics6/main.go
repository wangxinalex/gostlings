// Concept: two type parameters and inference at the call site
// Task: define a generic Transform function so this program compiles and runs
// Expected output: [2 4 6]
// [1 2 3]
// Hint: [T, U any] maps values of one type to another; Go infers both type
// arguments from the call, so main never writes them explicitly (Go Tour: Generics 1)

package main

import "fmt"

func Transform[T, U any](items []T, fn func(T) U) []U {
	transformed := make([]U, 0, len(items))
	for _, item := range items {
		transformed = append(transformed, fn(item))
	}
	return transformed
}

func main() {
	fmt.Println(Transform([]int{1, 2, 3}, func(v int) int { return v * 2 }))
	fmt.Println(Transform([]string{"a", "bb", "ccc"}, func(s string) int { return len(s) }))
}
