package planner

import "time"

func WaitUntilNoon() {
	now := time.Now()
	noon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())

	if now.After(noon) {
		noon = noon.Add(24 * time.Hour)
	}

	time.Sleep(time.Until(noon))
}
