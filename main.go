package main

import (
	"household-planner/pkg/household"
)

const (
	SENDER   = "+14155238886"
)

func main() {
	config := household.NewConfig()
	household.AssignTasks(config.DailyTasks, config.Members)

	client := household.InitializeTwilioClient()

	for _,member := range config.Members {
		assignedTasks := household.GetAssignedTasks(config.DailyTasks, member)
		dailyTaskMessage := household.CreateDailyTaskMessage(assignedTasks, member)
		household.SendMessage(client, dailyTaskMessage, SENDER, member.PhoneNumber)
	}
	
}
