// Concept: preserving lower-level error types while adding context
// Task: wrap strconv.Atoi's error in parsePort, then use errors.As in main
// Expected output: invalid port input
// Hint: return fmt.Errorf("parse port %q: %w", raw, err) from parsePort. In main,
//       declare var parseErr *strconv.NumError and pass &parseErr to errors.As;
//       checking the extracted type is more reliable than comparing error text.

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parsePort(raw string) (int, error) {
	port, err := strconv.Atoi(raw)
	if err != nil {
		// TODO: Add operation context while preserving the original error.
		return 0, nil
	}
	return port, nil
}

func main() {
	_, err := parsePort("eighty")
	var parseErr *strconv.NumError
	if errors.As(err, &parseErr) {
		// TODO: Keep this branch, which should run after parsePort preserves the cause.
		fmt.Println("invalid port input")
		return
	}
	fmt.Println("port parsed")
}
