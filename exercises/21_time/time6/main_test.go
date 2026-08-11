package main

import "testing"

func TestReuseTimerProducesOneEventPerGate(t *testing.T) {
	gates := make(chan struct{}, 3)
	for i := 0; i < 3; i++ {
		gates <- struct{}{}
	}
	close(gates)
	if got := reuseTimer(gates); got != 3 {
		t.Fatalf("reuseTimer() = %d, want 3", got)
	}
}
