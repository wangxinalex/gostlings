// Package lines counts logical lines in a string.
package lines

import "strings"

func Count(s string) int {
	if s == "" {
		return 0
	}
	return strings.Count(s, "\n")
}
