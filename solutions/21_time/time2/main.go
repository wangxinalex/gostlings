// Concept: time.Ticker for periodic tasks
// Task: the Ticker is created but never used; read from its channel in a select to print "tick" twice, then stop
// Expected output: tick
// tick
// Hint: t := time.NewTicker(d); its C field is a channel that fires every d; call t.Stop() to clean up (Go doc: time)

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for i := 0; i < 2; i++ {
		<-ticker.C
		fmt.Println("tick")
	}
}
