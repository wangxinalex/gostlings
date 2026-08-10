// Concept: composing validation, wrapping it at an operation boundary, and preserving causes
// Task: make createUser add operation context without hiding either validation error
// Expected output: create user: validation failed
// Hint: validateUser already returns errors.Join. In createUser, return
//       fmt.Errorf("create user: %w", err) when validation fails. In main,
//       use errors.Is twice; wrapping a joined error must still expose both causes.

package main

import (
	"errors"
	"fmt"
)

var (
	ErrUserNameRequired = errors.New("user name is required")
	ErrAgeInvalid       = errors.New("user age is invalid")
)

func validateUser(name string, age int) error {
	var errs []error
	if name == "" {
		errs = append(errs, ErrUserNameRequired)
	}
	if age < 0 {
		errs = append(errs, ErrAgeInvalid)
	}
	return errors.Join(errs...)
}

func createUser(name string, age int) error {
	err := validateUser(name, age)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func main() {
	err := createUser("", -1)
	if errors.Is(err, ErrUserNameRequired) && errors.Is(err, ErrAgeInvalid) {
		fmt.Println("create user: validation failed")
		return
	}
	fmt.Println("user created")
}
