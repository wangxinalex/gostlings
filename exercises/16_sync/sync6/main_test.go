package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestReadWriteCacheSupportsConcurrentReadersAndWriters(t *testing.T) {
	cache := &readWriteCache{}
	const writers = 8
	const readers = 16
	var wg sync.WaitGroup
	for i := 0; i < writers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			cache.Put(key, fmt.Sprintf("value-%d", i))
		}(i)
	}
	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i%writers)
			cache.Get(key)
		}(i)
	}
	wg.Wait()
	for i := 0; i < writers; i++ {
		key := fmt.Sprintf("key-%d", i)
		value, ok := cache.Get(key)
		if !ok || value != fmt.Sprintf("value-%d", i) {
			t.Fatalf("Get(%q) = (%q, %v), want matching value", key, value, ok)
		}
	}
}

func TestReadWriteCacheReportsMissingKeys(t *testing.T) {
	cache := &readWriteCache{}
	if value, ok := cache.Get("missing"); ok || value != "" {
		t.Fatalf("Get(missing) = (%q, %v), want empty false", value, ok)
	}
}
