package planner

import "time"

func WaitUntilNoon() {
	currentTime := time.Now()
	for currentTime.Hour() < 12 || currentTime.Hour() > 12 {
		time.Sleep(1 * time.Minute)
	}
}
