// Concept: sync.Once publishes one shared initialization.
// Task: make concurrent config callers observe one initialized value.
// Hint: put the assignment inside once.Do; return the published value afterward.
package main

import "sync"

var configOnce sync.Once
var configValue string
var initializeConfig = func() string { return "ready" }

func config() string {
	// TODO: Initialize config exactly once, even for concurrent callers.
	return ""
}

func main() {}
