// Concept: custom wrapper errors and the Unwrap contract
// Task: make StoreError preserve ErrPermission so the caller can classify it with errors.Is
// Expected output: save report: permission denied
// Hint: Error should describe the operation. Unwrap must return the embedded cause:
//       func (e *StoreError) Unwrap() error { return e.Err }. Then use errors.Is
//       in main; direct equality would fail once the cause is wrapped.

package main

import (
	"errors"
	"fmt"
)

var ErrPermission = errors.New("permission denied")

type StoreError struct {
	Op  string
	Err error
}

func (e *StoreError) Error() string {
	return e.Op + ": " + e.Err.Error()
}

func (e *StoreError) Unwrap() error {
	return e.Err
}

func saveReport() error {
	return &StoreError{Op: "save report", Err: ErrPermission}
}

func main() {
	err := saveReport()
	if errors.Is(err, ErrPermission) {
		fmt.Println(err)
		return
	}
	fmt.Println("report saved")
}
