// Concept: custom error types and errors.As for type extraction
// Task: use errors.As to extract the NotFoundError from the error and print the missing user ID
// Expected output: missing user ID: 42
// Hint: errors.As requires a pointer to a pointer: var e *NotFoundError; errors.As(err, &e) (Go Tour: Methods 20)

package main

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("user %d not found", e.ID)
}

func findUser(id int) error {
	return &NotFoundError{ID: 42}
}

func main() {
	err := findUser(42)
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		fmt.Println("missing user ID:", nfe.ID)
		return
	}
	fmt.Println("no error found")
}
