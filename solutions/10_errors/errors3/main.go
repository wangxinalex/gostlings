// Concept: sentinel errors, error wrapping with %w, and errors.Is
// Task: wrap ErrNotFound in the query function, then use errors.Is in main to check for it
// Expected output: user not found
// Hint: fmt.Errorf("query user: %w", ErrNotFound) wraps to preserve the sentinel (Go Tour: Methods 20)

package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("record not found")

func queryUser(id int) error {
	return fmt.Errorf("query user: %w", ErrNotFound)
}

func main() {
	err := queryUser(42)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("user not found")
		return
	}
	fmt.Println("everything is fine")
}
