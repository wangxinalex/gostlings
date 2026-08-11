// Concept: a reusable timer must be stopped and drained before Reset.
// Task: reset one timer for each gate event and count one timer event per gate.
// Hint: Stop, conditionally drain timer.C, Reset, then receive the new event.
package main

import "time"

func reuseTimer(gates <-chan struct{}) int {
	timer := time.NewTimer(time.Hour)
	defer timer.Stop()
	count := 0
	for range gates {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(0)
		<-timer.C
		count++
	}
	return count
}

func main() {}
