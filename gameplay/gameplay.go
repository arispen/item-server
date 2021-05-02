package gameplay

import (
	"fmt"
	"time"
)

func StartMainLoop() {
	// player attacks a monster
	//  - if hit then deal damage
	//	 - if killed then get exp, roll treasure, monster level ++
	// monster attacks the player character
	// 	- if hit then deal damage
	//	 - if player killed, go to monster level 1

	ticker := time.NewTicker(1000 * time.Millisecond)
	tocks := 0
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				tocks++
				if tocks%4 == 0 {
					fmt.Println("ias 0 at tock #", tocks)
				}
			}
		}
	}()

	time.Sleep(10 * 1000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
