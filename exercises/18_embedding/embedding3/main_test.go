package main

import "testing"

func TestEmployeeSatisfiesComposedInterface(t *testing.T) {
	var info EmployeeInfo = Employee{name: "Alice", id: 101}

	if got, want := info.Name(), "Alice"; got != want {
		t.Fatalf("Name() = %q, want %q", got, want)
	}
	if got, want := info.ID(), 101; got != want {
		t.Fatalf("ID() = %d, want %d", got, want)
	}
}
