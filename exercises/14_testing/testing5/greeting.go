// Package greeting writes a deterministic greeting to a file.
package greeting

import "os"

func WriteGreeting(path string) error {
	return os.WriteFile(path, []byte("hello\n"), 0o600)
}
