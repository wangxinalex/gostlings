package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestCancelChildDoesNotCancelParentOrSibling(t *testing.T) {
	parent := context.Background()
	sibling, cancelSibling := context.WithCancel(parent)
	returnedParent, child := cancelChild(parent)

	if returnedParent != parent {
		t.Fatal("cancelChild() did not return the supplied parent")
	}
	select {
	case <-child.Done():
	case <-time.After(500 * time.Millisecond):
		t.Fatal("cancelChild() did not cancel the child")
	}
	if !errors.Is(child.Err(), context.Canceled) {
		t.Fatalf("child.Err() = %v, want context.Canceled", child.Err())
	}
	if parent.Err() != nil {
		t.Fatalf("parent.Err() = %v, want nil after child cancellation", parent.Err())
	}
	select {
	case <-sibling.Done():
		t.Fatal("cancelChild() canceled a sibling")
	default:
	}

	cancelSibling()
	select {
	case <-sibling.Done():
	case <-time.After(500 * time.Millisecond):
		t.Fatal("sibling cancel function did not clean up the sibling")
	}
}
