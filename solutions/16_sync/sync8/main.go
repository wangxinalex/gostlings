// Concept: sync.Once publishes one shared initialization.
// Task: make concurrent config callers observe one initialized value.
// Hint: put the assignment inside once.Do; return the published value afterward.
package main

import "sync"

var configOnce sync.Once
var configValue string
var initializeConfig = func() string { return "ready" }

func config() string {
	configOnce.Do(func() { configValue = initializeConfig() })
	return configValue
}

func main() {}
