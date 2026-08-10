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
	// TODO: Return a wrapped error using fmt.Errorf with %w so errors.Is can detect ErrNotFound.
	return fmt.Errorf("query user: %v", ErrNotFound)
}

func main() {
	err := queryUser(42)
	// TODO: Use errors.Is to check if err wraps ErrNotFound, then print "user not found".
	if errors.Is(err, ErrNotFound) {
		fmt.Println("user not found")
		return
	}
	fmt.Println("everything is fine")
}
