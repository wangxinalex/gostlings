// Concept: translating internal errors at an application boundary
// Task: classify a wrapped not-found error and return a stable user-facing message
// Expected output: 404 user not found
// Hint: keep the internal context from findUser, but do not compare err.Error().
//       In lookupMessage, use errors.Is(err, ErrUserNotFound) for the 404 branch;
//       unknown errors should return the generic 500 message.

package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

func findUser(id int) error {
	return fmt.Errorf("query user %d: %w", id, ErrUserNotFound)
}

func lookupMessage(id int) string {
	err := findUser(id)
	if errors.Is(err, ErrUserNotFound) {
		return "404 user not found"
	}
	return "500 internal error"
}

func main() {
	fmt.Println(lookupMessage(7))
}
