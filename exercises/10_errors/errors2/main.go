// Concept: validating input and returning meaningful errors
// Task: the checkAge function returns nil even when age is negative; fix it so age -1 prints an error
// Expected output: age must not be negative
// Hint: return a non-nil error from the check so the condition matters (Go Tour: Methods 19-20)

package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {
	if age < 0 {
		// TODO: Return an error instead of nil.
		return errors.New("age must be non-negative")
	}
	return nil
}

func main() {
	err := checkAge(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("age is valid")
}
