// Concept: sorting slices with sort.Slice
// Task: sort the numbers slice in descending order using sort.Slice
// Expected output: [9 7 5 3 1]
// Hint: functions6 introduced function values. sort.Slice(s, func(i, j int) bool
//       { return s[i] > s[j] }) sorts descending (Go doc: sort)

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 1, 7, 9, 5}
	// TODO: Use sort.Slice to sort nums in descending order.
	fmt.Println(nums)
}
