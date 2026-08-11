package main

import (
	"sync"
	"testing"
)

func TestRegistryStoresAndLoadsValues(t *testing.T) {
	var r registry
	r.Store("a", "alpha")
	if value, ok := r.Load("a"); !ok || value != "alpha" {
		t.Fatalf("Load(a) = (%q, %v), want alpha", value, ok)
	}
	if _, ok := r.Load("missing"); ok {
		t.Fatal("Load(missing) reported a value")
	}
}

func TestRegistryLoadOrStoreKeepsOneCanonicalValue(t *testing.T) {
	var r registry
	const callers = 20
	values := make(chan string, callers)
	var wg sync.WaitGroup
	for i := 0; i < callers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, _ := r.LoadOrStore("shared", "candidate")
			values <- value
		}(i)
	}
	wg.Wait()
	close(values)
	for value := range values {
		if value != "candidate" {
			t.Fatalf("canonical value = %q, want candidate", value)
		}
	}
}
