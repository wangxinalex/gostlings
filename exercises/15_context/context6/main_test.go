package main

import (
	"context"
	"testing"
	"time"
)

func TestStartChildrenPropagatesParentCancellationToEveryChild(t *testing.T) {
	const count = 3
	previous := childStopped
	stopped := make(chan struct{}, count)
	childStopped = func() { stopped <- struct{}{} }
	previousChildStarted := childStarted
	started := make(chan struct{}, count)
	childStarted = func() { started <- struct{}{} }
	previousWithCancel := withCancel
	canceled := make(chan struct{}, count)
	release := make(chan struct{})
	released := false
	releaseCallbacks := func() {
		if !released {
			close(release)
			released = true
		}
	}
	withCancel = func(parent context.Context) (context.Context, context.CancelFunc) {
		ctx, cancel := context.WithCancel(parent)
		return ctx, func() {
			canceled <- struct{}{}
			<-release
			cancel()
		}
	}
	t.Cleanup(func() {
		childStopped = previous
		childStarted = previousChildStarted
		withCancel = previousWithCancel
	})

	parent, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
		releaseCallbacks()
	})
	done := startChildren(parent, count)
	for child := 0; child < count; child++ {
		select {
		case <-started:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("child %d was not created before parent cancellation", child)
		}
	}
	cancel()

	for child := 0; child < count; child++ {
		select {
		case <-stopped:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("child %d did not observe parent cancellation", child)
		}
	}
	for child := 0; child < count; child++ {
		select {
		case <-canceled:
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("child %d did not call its cancel function", child)
		}
	}
	select {
	case <-done:
		t.Fatal("startChildren() closed done before child cancel functions returned")
	default:
	}

	releaseCallbacks()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("startChildren() did not close done after every child stopped")
	}
}

func TestStartChildrenWithNoChildrenClosesDone(t *testing.T) {
	select {
	case <-startChildren(context.Background(), 0):
	case <-time.After(500 * time.Millisecond):
		t.Fatal("startChildren() did not close done for an empty group")
	}
}
