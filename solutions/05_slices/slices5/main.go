// Concept: sorting slices with sort.Slice
// Task: sort the numbers slice in descending order using sort.Slice
// Expected output: [9 7 5 3 1]
// Hint: sort.Slice(s, func(i, j int) bool { return s[i] > s[j] }) sorts descending (Go doc: sort)

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 1, 7, 9, 5}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)
}
