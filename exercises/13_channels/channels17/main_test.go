package main

import (
	"reflect"
	"testing"
	"time"
)

func TestRunOrderedRestoresInputOrder(t *testing.T) {
	result := make(chan []int, 1)
	go func() { result <- runOrdered(4, []int{1, 2, 3, 4, 5}) }()

	select {
	case got := <-result:
		want := []int{1, 4, 9, 16, 25}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("runOrdered() = %v, want %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("runOrdered() did not finish")
	}
}
