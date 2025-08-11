package main

import (
	"fmt"
	"household-planner/pkg/backend"
	"household-planner/pkg/planner"
)

func main() {
	fmt.Println("[INFO] Starting Household Planner...")
	go backend.StartServer()

	config := planner.LoadConfig()
	myHousehold := planner.NewHousehold(config)

	currentMemberIndex := 0
	for {
		currentMember := myHousehold.Members[currentMemberIndex]
		currentMemberIndex++
		if currentMemberIndex >= len(myHousehold.Members) {
			currentMemberIndex = 0
		}

		myHousehold.ClearAssignments()
		myHousehold.AssignDailyTasks()
		myHousehold.AssignWeeklyTasks(currentMember)
		myHousehold.AssignMonthlyTasks(currentMember)

		client := planner.InitializeTwilioClient()
		for _, member := range myHousehold.Members {
			assignedTasks := myHousehold.GetAssignedTasks(member)
			dailyTaskMessage := planner.CreateDailyTaskMessage(assignedTasks, member)
			planner.SendMessage(client, dailyTaskMessage, member.Phonenumber)
		}

		planner.WaitUntilNoon()
	}
}
