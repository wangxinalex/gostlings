// Concept: counting with range, if/else if, and named results
// Task: complete classify so a single pass counts every value as negative, even,
//       or odd; a negative number counts only as negative
// Expected output: evens=3 odds=2 negatives=3
// Hint: the named results evens, odds and negatives are already declared by the
//       signature. range yields each value, and an if / else if / else chain
//       keeps the three counts exclusive (Go Tour: Flowcontrol 1-3)

package main

import "fmt"

func classify(values []int) (evens, odds, negatives int) {
	for _, v := range values {
		if v < 0 {
			negatives++
		} else if v%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return evens, odds, negatives
}

func main() {
	values := []int{4, -3, 0, 7, -8, 12, -5, 9}
	evens, odds, negatives := classify(values)
	fmt.Printf("evens=%d odds=%d negatives=%d\n", evens, odds, negatives)
}
