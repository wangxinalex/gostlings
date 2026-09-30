package main

import (
	"context"
	"errors"
	"slices"
	"testing"
	"time"
)

func TestGatherStartsEveryTaskConcurrentlyAndKeepsOrder(t *testing.T) {
	const count = 4
	started := make(chan struct{}, count)
	gates := make([]chan struct{}, count)
	tasks := make([]func(context.Context) (int, error), count)
	for index := range tasks {
		gates[index] = make(chan struct{})
		tasks[index] = func(ctx context.Context) (int, error) {
			started <- struct{}{}
			select {
			case <-gates[index]:
				return index * 10, nil
			case <-ctx.Done():
				return 0, ctx.Err()
			}
		}
	}
	type outcome struct {
		values []int
		err    error
	}
	done := make(chan outcome, 1)
	go func() {
		values, err := gather(context.Background(), tasks)
		done <- outcome{values: values, err: err}
	}()

	for i := 0; i < count; i++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatal("gather() did not start every task concurrently")
		}
	}
	for index := count - 1; index >= 0; index-- {
		close(gates[index])
	}

	select {
	case got := <-done:
		if got.err != nil {
			t.Fatalf("gather() error = %v, want nil", got.err)
		}
		want := []int{0, 10, 20, 30}
		if !slices.Equal(got.values, want) {
			t.Fatalf("gather() = %v, want %v in task order", got.values, want)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("gather() did not return after every task finished")
	}
}

func TestGatherCancelsAndJoinsEveryTaskAfterTheFirstError(t *testing.T) {
	failure := errors.New("task 2 failed")
	const count = 4
	stopped := make(chan int, count)
	tasks := make([]func(context.Context) (int, error), count)
	for index := range tasks {
		if index == 2 {
			tasks[index] = func(context.Context) (int, error) { return 0, failure }
			continue
		}
		tasks[index] = func(ctx context.Context) (int, error) {
			<-ctx.Done()
			stopped <- index
			return 0, ctx.Err()
		}
	}

	_, err := gather(context.Background(), tasks)
	if !errors.Is(err, failure) {
		t.Fatalf("gather() error = %v, want %v", err, failure)
	}
	for i := 0; i < count-1; i++ {
		select {
		case <-stopped:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("gather() returned before joining goroutine %d", i)
		}
	}
	select {
	case index := <-stopped:
		t.Fatalf("gather() ran task %d more than once", index)
	default:
	}
}

func TestGatherDoesNotStartWorkWithACanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := make(chan struct{}, 1)
	tasks := []func(context.Context) (int, error){
		func(context.Context) (int, error) {
			started <- struct{}{}
			return 1, nil
		},
	}

	values, err := gather(ctx, tasks)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("gather() error = %v, want context.Canceled", err)
	}
	if len(values) != 0 {
		t.Fatalf("gather() values = %v, want none", values)
	}
	select {
	case <-started:
		t.Fatal("gather() started a task with an already canceled context")
	default:
	}
}

func TestGatherWithoutTasksReturnsAnEmptySlice(t *testing.T) {
	values, err := gather(context.Background(), nil)
	if err != nil {
		t.Fatalf("gather() error = %v, want nil", err)
	}
	if values == nil {
		t.Fatal("gather() = nil, want an empty non-nil slice")
	}
	if len(values) != 0 {
		t.Fatalf("gather() values = %v, want empty", values)
	}
}
