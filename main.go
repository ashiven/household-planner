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

	config := planner.LoadConfig()
	myHousehold := planner.NewHousehold(config)

	backend.SetConfig(config)
	go backend.StartServer()

	for {
		fmt.Println("[INFO] A new day has started, assigning tasks...")

		myHousehold.ClearAssignments()
		myHousehold.UpdateCurrentMember()
		myHousehold.AssignDailyTasks()
		myHousehold.AssignWeeklyTasks()
		myHousehold.AssignMonthlyTasks()

		client := planner.InitializeTwilioClient()
		for _, member := range *myHousehold.Members {
			assignedTasks := myHousehold.GetAssignedTasks(member)
			planner.SendMessageSms(client, member, assignedTasks, debug)
		}

		if debug {
			fmt.Println("[DEBUG] Household members: ")
			fmt.Println(myHousehold.Members)
			fmt.Println("[DEBUG] Household Daily tasks: ")
			fmt.Println(myHousehold.DailyTasks)
			fmt.Println("[DEBUG] Household Weekly tasks: ")
			fmt.Println(myHousehold.WeeklyTasks)
			fmt.Println("[DEBUG] Household Monthly tasks: ")
			fmt.Println(myHousehold.MonthlyTasks)
			fmt.Println("[DEBUG] Config members: ")
			fmt.Println(config.Members)
			fmt.Println("[DEBUG] Config Daily tasks: ")
			fmt.Println(config.DailyTasks)
			fmt.Println("[DEBUG] Config Weekly tasks: ")
			fmt.Println(config.WeeklyTasks)
			fmt.Println("[DEBUG] Config Monthly tasks: ")
			fmt.Println(config.MonthlyTasks)
			time.Sleep(2 * time.Minute)
		} else {
			planner.WaitUntilNoon()
		}
	}
}
