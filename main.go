package main

import (
	"fmt"
	"household-planner/pkg/backend"
	"household-planner/pkg/planner"
	"os"
	"time"
)

func main() {
	fmt.Println("[INFO] Starting Household Planner...")
	debug := len(os.Args) > 1 && os.Args[1] == "-d"

	myHousehold, err := planner.NewHousehold()
	if err != nil {
		fmt.Println("[ERROR] Failed to create household:", err)
		return
	}

	backend.SetHousehold(myHousehold)
	go backend.StartServer()

	for {
		fmt.Println("[INFO] A new day has started, assigning tasks...")

		myHousehold.ClearAssignments()
		myHousehold.UpdateCurrentMember()
		myHousehold.AssignDailyTasks()
		myHousehold.AssignWeeklyTasks()
		myHousehold.AssignMonthlyTasks()

		client := planner.InitializeTwilioClient()
		for _, member := range myHousehold.Members {
			assignedTasks := myHousehold.GetAssignedTasks(member)
			planner.SendMessageSms(client, member, assignedTasks, debug)
		}

		if debug {
			fmt.Println("[DEBUG] Starting next day in two minutes...: ")
			time.Sleep(2 * time.Minute)
		} else {
			planner.WaitUntilNoon()
		}
	}
}
