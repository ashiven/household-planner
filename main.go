package main

import (
	"household-planner/pkg/household"
)

func main() {
	config := household.NewConfig()
	weeklyTasksPerDay := len(config.WeeklyTasks) / len(config.Members)

	currentMemberIndex := 0
	for {
		currentMember := config.Members[currentMemberIndex]

		household.AssignTasksToAll(config.DailyTasks, config.Members)
		household.AssignTasks(config.WeeklyTasks, currentMember, weeklyTasksPerDay)
		// TODO:
		// household.AssignTasks(config.MonthlyTasks, currentMember)

		client := household.InitializeTwilioClient()
		for _, member := range config.Members {
			assignedTasks := household.GetAssignedTasks(config, member)
			dailyTaskMessage := household.CreateDailyTaskMessage(assignedTasks, member)
			household.SendMessage(client, dailyTaskMessage, member.PhoneNumber)
		}

		household.ClearTasks(config.DailyTasks)
		household.ClearTasks(config.WeeklyTasks)
		household.ClearTasks(config.MonthlyTasks)

		currentMemberIndex++
		if currentMemberIndex >= len(config.Members) {
			currentMemberIndex = 0
		}

		household.WaitUntilNoon()
	}
}
