package main

import (
	"sync"
	"testing"
)

func TestConfigInitializesExactlyOnce(t *testing.T) {
	configOnce = sync.Once{}
	configValue = ""
	previous := initializeConfig
	var mu sync.Mutex
	loads := 0
	initializeConfig = func() string {
		mu.Lock()
		loads++
		mu.Unlock()
		return "loaded-once"
	}
	t.Cleanup(func() { initializeConfig = previous })

	const callers = 20
	values := make(chan string, callers)
	var wg sync.WaitGroup
	for i := 0; i < callers; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); values <- config() }()
	}
	wg.Wait()
	close(values)
	for value := range values {
		if value != "loaded-once" {
			t.Fatalf("config() = %q, want loaded-once", value)
		}
	}
	if loads != 1 {
		t.Fatalf("initializer called %d times, want once", loads)
	}
}
