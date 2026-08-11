// Concept: sync.Once can publish context-aware configuration exactly once.
// Task: call load once, pass the caller context unchanged, and return one stable value.
// Hint: put the load call inside Once.Do and keep the stored value protected by that publication.
package main

import (
	"context"
	"sync"
)

var configOnce sync.Once
var configValue string

func loadConfig(ctx context.Context, load func(context.Context) string) string {
	configOnce.Do(func() { configValue = load(ctx) })
	return configValue
}

func main() {}
