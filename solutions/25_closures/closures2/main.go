// Concept: functions are values — passing behavior as an argument
// Task: implement transform so it returns a new slice with fn applied to every element
// Expected output: [4 9 16]
// Hint: allocate the result with make([]int, len(xs)) and set out[i] = fn(xs[i]);
//       the parameter fn is just another value you can call
//       (builds on Go Tour: Basics 4-7; function values are not covered in the Tour)

package main

import "fmt"

func transform(xs []int, fn func(int) int) []int {
	out := make([]int, len(xs))
	for i, x := range xs {
		out[i] = fn(x)
	}
	return out
}

func main() {
	fmt.Println(transform([]int{2, 3, 4}, func(n int) int { return n * n }))
}
