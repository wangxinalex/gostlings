// Concept: sorting built-in slices with sort.Strings and sort.Ints
// Task: complete sortNames and sortNumbers so they sort the input in place
// Expected output: [ada bob eva]
// [1 2 3 7]
// Hint: the sort package has ready-made sorts for the built-in ordered types, and
//       they reorder the slice you pass instead of returning a new one (Go doc: sort)
// Stuck?: sort.Strings and sort.Ints; there is no return value to collect.

package main

import (
	"fmt"
	"sort"
)

func sortNames(names []string) []string {
	// TODO: Sort names in place and return it.
	return names
}

func sortNumbers(nums []int) []int {
	// TODO: Sort nums in place and return it.
	return nums
}

func main() {
	fmt.Println(sortNames([]string{"bob", "ada", "eva"}))
	fmt.Println(sortNumbers([]int{7, 1, 3, 2}))
}
