// Concept: a reusable timer must be stopped and drained before Reset.
// Task: reset one timer for each gate event and count one timer event per gate.
// Hint: Stop, conditionally drain timer.C, Reset, then receive the new event.
package main

func reuseTimer(gates <-chan struct{}) int {
	// TODO: Reuse one timer safely and count one event for every gate.
	return 0
}

func main() {}
