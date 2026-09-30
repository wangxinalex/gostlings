// Concept: time.Ticker for periodic tasks
// Task: the Ticker is created but never used; read from its channel in a select to print "tick" twice, then stop
// Expected output: tick
// tick
// Hint: a ticker exposes a channel that delivers one value per interval; receive
//       twice from it and stop the ticker before returning (Go doc: time)
// Stuck?: time.NewTicker's channel field is C, and Stop releases the ticker.

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(50 * time.Millisecond)
	// TODO: Use a timer or count to stop after 2 ticks.
	//       Print "tick" each time the ticker fires.

	time.Sleep(200 * time.Millisecond)
	ticker.Stop()
	fmt.Println("never ticked")
}
