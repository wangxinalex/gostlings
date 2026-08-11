package main

import (
	"reflect"
	"testing"
	"time"
)

func TestGenerateAndSumCompleteAfterClose(t *testing.T) {
	ch := generate(1, 2, 3)
	result := make(chan int, 1)
	go func() { result <- sum(ch) }()
	select {
	case got := <-result:
		if got != 6 {
			t.Fatalf("sum() = %d, want 6", got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("sum() did not finish after producer completion")
	}
	var values []int
	for value := range generate(4, 5) {
		values = append(values, value)
	}
	if !reflect.DeepEqual(values, []int{4, 5}) {
		t.Fatalf("generated values = %v, want [4 5]", values)
	}
}
