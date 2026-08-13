// Concept: sorting built-in slices with sort.Strings and sort.Ints
// Task: complete sortNames and sortNumbers so they sort the input in place
// Expected output: [ada bob eva]
// [1 2 3 7]
// Hint: sort.Strings(s) and sort.Ints(s) sort in place (Go doc: sort)

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
