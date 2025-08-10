package main

import (
	"household-planner/pkg/household"
)

func main() {
	config := household.NewConfig()

	currentMemberIndex := 0
	for {
		currentMember := config.Members[currentMemberIndex]

		household.AssignTasksToAll(config.DailyTasks, config.Members)
		// maybe pop them from the list in equal weekly shares
		// and when the list is empty, refill it from the config
		household.AssignTasks(config.WeeklyTasks, currentMember)
		household.AssignTasks(config.MonthlyTasks, currentMember)

		client := household.InitializeTwilioClient()
		for _, member := range config.Members {
			assignedTasks := household.GetAssignedTasks(config.DailyTasks, member)
			dailyTaskMessage := household.CreateDailyTaskMessage(assignedTasks, member)
			household.SendMessage(client, dailyTaskMessage, member.PhoneNumber)
		}

		currentMemberIndex++
		if currentMemberIndex >= len(config.Members) {
			currentMemberIndex = 0
		}

		household.WaitUntilNoon()
	}
}
