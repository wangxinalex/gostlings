package main

import (
	"context"
	"sync"
	"testing"
)

func TestLoadConfigPassesContextAndInitializesOnce(t *testing.T) {
	configOnce = sync.Once{}
	configValue = ""
	ctx := context.WithValue(context.Background(), struct{}{}, "request")
	calls := 0
	var mu sync.Mutex
	load := func(got context.Context) string {
		if got != ctx {
			t.Errorf("load received a different context")
		}
		mu.Lock()
		calls++
		mu.Unlock()
		return "loaded"
	}
	values := make(chan string, 2)
	go func() { values <- loadConfig(ctx, load) }()
	go func() { values <- loadConfig(ctx, load) }()
	if <-values != "loaded" || <-values != "loaded" {
		t.Fatal("callers did not see one config")
	}
	if calls != 1 {
		t.Fatalf("load called %d times, want once", calls)
	}
}
