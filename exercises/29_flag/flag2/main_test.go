package main

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	port, rest, err := parseCommand([]string{"-port", "8080", "alpha", "beta"})
	if err != nil {
		t.Fatal(err)
	}
	if port != 8080 {
		t.Fatalf("parseCommand() port = %d, want 8080", port)
	}
	if want := []string{"alpha", "beta"}; !reflect.DeepEqual(rest, want) {
		t.Fatalf("parseCommand() rest = %v, want %v", rest, want)
	}
}

func TestParseCommandRejectsUnknownFlag(t *testing.T) {
	if _, _, err := parseCommand([]string{"-unknown"}); err == nil {
		t.Fatal("parseCommand() accepted an unknown flag")
	}
}
