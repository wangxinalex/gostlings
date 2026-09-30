// Concept: labelled loops with break
// Task: complete firstCommon so it returns the first value of left that also
//       appears in right, or -1 when there is none
// Expected output: first common: 16
// no common: -1
// Hint: label the outer loop (outer:) and use `break outer` to leave both loops
//       as soon as a match is found; a plain break would only end the inner
//       loop, and `continue outer` would move on to the next left value
//       (Go Tour: Flowcontrol 8)

package main

import "fmt"

func firstCommon(left, right []int) int {
	result := -1
outer:
	for _, l := range left {
		for _, r := range right {
			if l == r {
				result = l
				break outer
			}
		}
	}
	return result
}

func main() {
	left := []int{4, 8, 15, 16, 23, 42}
	fmt.Println("first common:", firstCommon(left, []int{1, 2, 3, 16, 23}))
	fmt.Println("no common:", firstCommon(left, []int{7, 11, 13}))
}
