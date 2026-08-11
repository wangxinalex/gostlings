package main

import (
	"reflect"
	"testing"
	"time"
)

func TestParallelBoundsActiveWorkAndRestoresOrder(t *testing.T) {
	started := make(chan int, 2)
	activeSlots := make(chan struct{}, 2)
	activeSlots <- struct{}{}
	activeSlots <- struct{}{}
	overflow := make(chan int, 1)
	release := make(chan struct{})
	returned := make(chan []int, 1)
	go func() {
		returned <- parallel(2, []int{3, 1, 2}, func(value int) int {
			select {
			case <-activeSlots:
				started <- value
			default:
				select {
				case <-release:
				default:
					overflow <- value
				}
			}
			<-release
			return value * value
		})
	}()

	for worker := 0; worker < 2; worker++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("parallel() did not start work up to its limit")
		}
	}
	close(release)
	select {
	case got := <-returned:
		if want := []int{9, 1, 4}; !reflect.DeepEqual(got, want) {
			t.Fatalf("parallel() = %v, want ordered %v", got, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("parallel() did not finish after work was released")
	}
	select {
	case value := <-overflow:
		t.Fatalf("parallel() started a third active job (%d) with limit 2", value)
	default:
	}
}
