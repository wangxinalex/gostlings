// Concept: collecting independent validation errors with errors.Join
// Task: return both validation failures from validateConfig and detect each cause in main
// Expected output: configuration has 2 errors
// Hint: collect non-nil errors in a slice, then return errors.Join(errs...).
//       errors.Is works through a joined error, so check ErrNameRequired and
//       ErrPortRange separately instead of parsing the combined error string.

package main

import (
	"errors"
	"fmt"
)

var (
	ErrNameRequired = errors.New("name is required")
	ErrPortRange    = errors.New("port is out of range")
)

func validateConfig(name string, port int) error {
	var errs []error
	if name == "" {
		errs = append(errs, ErrNameRequired)
	}
	if port < 1 || port > 65535 {
		errs = append(errs, ErrPortRange)
	}
	return errors.Join(errs...)
}

func main() {
	err := validateConfig("", 70000)
	if errors.Is(err, ErrNameRequired) && errors.Is(err, ErrPortRange) {
		fmt.Println("configuration has 2 errors")
		return
	}
	fmt.Println("configuration is valid")
}
