package planner

import (
	"math/rand"
)

func (household *Household) UpdateCurrentMember() {
	household.currentMember = household.Members[household.currentMemberIndex]
	household.currentMemberIndex++
	if household.currentMemberIndex >= len(household.Members) {
		household.currentMemberIndex = 0
	}
}

func (household *Household) AssignDailyTasks() {
	rand.Shuffle(len(household.DailyTasks), func(i, j int) {
		household.DailyTasks[i], household.DailyTasks[j] = household.DailyTasks[j], household.DailyTasks[i]
	})

	shuffledMembers := make([]*Member, len(household.Members))
	copy(shuffledMembers, (household.Members))
	rand.Shuffle(len(shuffledMembers), func(i, j int) {
		shuffledMembers[i], shuffledMembers[j] = shuffledMembers[j], shuffledMembers[i]
	})

	assigneeIndex := 0
	for _, task := range household.DailyTasks {
		task.SetAssignee(shuffledMembers[assigneeIndex])
		assigneeIndex++
		if assigneeIndex >= len(shuffledMembers) {
			assigneeIndex = 0
		}
	}
}

func (household *Household) AssignWeeklyTasks() {
	if household.remainingWeeklyTasks == 0 {
		household.remainingWeeklyTasks = len(household.WeeklyTasks)
	}

	amountAdded := 0
	weeklyTasksPerDay := max(len(household.WeeklyTasks)/len(household.Members), 1)
	for amountAdded < weeklyTasksPerDay && household.remainingWeeklyTasks > 0 {
		currentTaskIndex := len(household.WeeklyTasks) - household.remainingWeeklyTasks
		task := household.WeeklyTasks[currentTaskIndex]
		task.SetAssignee(household.currentMember)

		household.remainingWeeklyTasks--
		amountAdded++
	}
}

func (household *Household) AssignMonthlyTasks() {
	if household.remainingMonthlyTasks == 0 {
		household.remainingMonthlyTasks = len(household.MonthlyTasks)
	}

	randomMember := household.Members[rand.Intn(len(household.Members))]
	for randomMember.Name == household.currentMember.Name {
		randomMember = household.Members[rand.Intn(len(household.Members))]
	}

	taskIntervalMonth := 30 / len(household.MonthlyTasks)
	if household.dayOfTheMonth%taskIntervalMonth == 0 && household.remainingMonthlyTasks > 0 {
		currentTaskIndex := len(household.MonthlyTasks) - household.remainingMonthlyTasks
		task := household.MonthlyTasks[currentTaskIndex]
		task.SetAssignee(randomMember)

		household.remainingMonthlyTasks--
	}

	household.dayOfTheMonth++
	if household.dayOfTheMonth > 30 {
		household.dayOfTheMonth = 1
	}
}

func (household *Household) ClearAssignments() {
	for _, task := range household.DailyTasks {
		task.SetAssignee(nil)
	}
	for _, task := range household.WeeklyTasks {
		task.SetAssignee(nil)
	}
	for _, task := range household.MonthlyTasks {
		task.SetAssignee(nil)
	}
}

func (household *Household) GetAssignedTasks(member *Member) []Assignable {
	assignedTasks := []Assignable{}

	for _, task := range household.DailyTasks {
		assignee := task.GetAssignee()
		if assignee != nil && assignee.Name == member.Name {
			assignedTasks = append(assignedTasks, task)
		}
	}

	for _, task := range household.WeeklyTasks {
		assignee := task.GetAssignee()
		if assignee != nil && assignee.Name == member.Name {
			assignedTasks = append(assignedTasks, task)
		}
	}

	for _, task := range household.MonthlyTasks {
		assignee := task.GetAssignee()
		if assignee != nil && assignee.Name == member.Name {
			assignedTasks = append(assignedTasks, task)
		}
	}

	return assignedTasks
}
