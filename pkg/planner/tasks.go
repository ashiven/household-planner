// Package planner provides functionality to manage household members.
//
// It includes the ability to create a new household member with a name and phone number.
package planner

import (
	"math/rand"
)

func (household *Household) AssignDailyTasks() {
	rand.Shuffle(len(household.DailyTasks), func(i, j int) {
		household.DailyTasks[i], household.DailyTasks[j] = household.DailyTasks[j], household.DailyTasks[i]
	})

	shuffledMembers := make([]*Member, len(household.Members))
	copy(shuffledMembers, household.Members)
	rand.Shuffle(len(shuffledMembers), func(i, j int) {
		shuffledMembers[i], shuffledMembers[j] = shuffledMembers[j], shuffledMembers[i]
	})

	currentMemberIndex := 0
	for _, task := range household.DailyTasks {
		task.SetAssignee(shuffledMembers[currentMemberIndex])
		currentMemberIndex++
		if currentMemberIndex >= len(shuffledMembers) {
			currentMemberIndex = 0
		}
	}
}

func (household *Household) AssignWeeklyTasks(member *Member) {
	if len(household.remainingWeeklyTasks) == 0 {
		household.remainingWeeklyTasks = append(household.remainingWeeklyTasks, household.WeeklyTasks...)
	}

	amountAdded := 0
	for amountAdded < household.weeklyTasksPerDay && len(household.remainingWeeklyTasks) > 0 {
		task := household.remainingWeeklyTasks[0]
		task.SetAssignee(member)
		household.remainingWeeklyTasks = household.remainingWeeklyTasks[1:]
		amountAdded++
	}
}

func (household *Household) AssignMonthlyTasks(member *Member) {
	if len(household.remainingMonthlyTasks) == 0 {
		household.remainingMonthlyTasks = append(household.remainingMonthlyTasks, household.MonthlyTasks...)
	}

	randomMember := household.Members[rand.Intn(len(household.Members))]
	for randomMember.Name == member.Name {
		randomMember = household.Members[rand.Intn(len(household.Members))]
	}

	if household.dayOfTheMonth%household.taskIntervalMonth == 0 {
		task := household.remainingMonthlyTasks[0]
		task.SetAssignee(randomMember)
		household.remainingMonthlyTasks = household.remainingMonthlyTasks[1:]
	}

	household.dayOfTheMonth++
	if household.dayOfTheMonth > 30 {
		household.dayOfTheMonth = 1
	}
}

func (household *Household) ClearTasks() {
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
