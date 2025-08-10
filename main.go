package main

import (
	"household-planner/pkg/planner"
)

func main() {
	config := planner.LoadConfig()
	myHousehold := planner.NewHousehold(config)

	currentMemberIndex := 0
	for {
		currentMember := myHousehold.Members[currentMemberIndex]
		currentMemberIndex++
		if currentMemberIndex >= len(myHousehold.Members) {
			currentMemberIndex = 0
		}

		myHousehold.AssignDailyTasks()
		myHousehold.AssignWeeklyTasks(currentMember)
		myHousehold.AssignMonthlyTasks(currentMember)

		client := planner.InitializeTwilioClient()
		for _, member := range myHousehold.Members {
			assignedTasks := myHousehold.GetAssignedTasks(member)
			dailyTaskMessage := planner.CreateDailyTaskMessage(assignedTasks, member)
			planner.SendMessage(client, dailyTaskMessage, member.PhoneNumber)
		}

		myHousehold.ClearTasks()
		planner.WaitUntilNoon()
	}
}
