package planner

import "math/rand"

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
	if len(household.remainingWeeklyTasks) == 0 {
		household.remainingWeeklyTasks = append(household.remainingWeeklyTasks, (household.WeeklyTasks)...)
	}

	amountAdded := 0
	weeklyTasksPerDay := max(len(household.WeeklyTasks)/len(household.Members), 1)
	for amountAdded < weeklyTasksPerDay && len(household.remainingWeeklyTasks) > 0 {
		task := household.remainingWeeklyTasks[0]
		task.SetAssignee(household.currentMember)
		household.remainingWeeklyTasks = household.remainingWeeklyTasks[1:]
		amountAdded++
	}
}

func (household *Household) AssignMonthlyTasks() {
	if len(household.remainingMonthlyTasks) == 0 {
		household.remainingMonthlyTasks = append(household.remainingMonthlyTasks, household.MonthlyTasks...)
	}

	randomMember := household.Members[rand.Intn(len(household.Members))]
	for randomMember.Name == household.currentMember.Name {
		randomMember = household.Members[rand.Intn(len(household.Members))]
	}

	taskIntervalMonth := 30 / len(household.MonthlyTasks)
	if household.dayOfTheMonth%taskIntervalMonth == 0 {
		task := household.remainingMonthlyTasks[0]
		task.SetAssignee(randomMember)
		household.remainingMonthlyTasks = household.remainingMonthlyTasks[1:]
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
