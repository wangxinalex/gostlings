// Concept: functions are values — passing behavior as an argument
// Task: implement transform so it returns a new slice with fn applied to every element
// Expected output: [4 9 16]
// Hint: allocate an output slice as long as the input, then fill each position by
//       calling the function value you were given with the matching input element
//       (builds on Go Tour: Basics 4-7; function values are not covered in the Tour)

package main

import "fmt"

func transform(xs []int, fn func(int) int) []int {
	// TODO: Return a new slice containing fn(x) for each x in xs.
}

func main() {
	fmt.Println(transform([]int{2, 3, 4}, func(n int) int { return n * n }))
}
