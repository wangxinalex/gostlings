// Package lookup exposes a classified not-found error.
package lookup

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("record not found")

func Find(id int) (string, error) {
	return "", fmt.Errorf("find %d: %w", id, ErrNotFound)
}
