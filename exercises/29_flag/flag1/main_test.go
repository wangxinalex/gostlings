package main

import "testing"

func TestParseArgs(t *testing.T) {
	name, count, verbose, err := parseArgs([]string{"-name", "Ada", "-count", "3", "-verbose"})
	if err != nil {
		t.Fatal(err)
	}
	if name != "Ada" || count != 3 || !verbose {
		t.Fatalf("parseArgs() = (%q, %d, %t), want (Ada, 3, true)", name, count, verbose)
	}
}

func TestParseArgsRejectsUnknownFlag(t *testing.T) {
	if _, _, _, err := parseArgs([]string{"-unknown"}); err == nil {
		t.Fatal("parseArgs() accepted an unknown flag")
	}
}
