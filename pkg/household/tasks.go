// Package household provides functionality to manage household members.
//
// It includes the ability to create a new household member with a name and phone number.
package household

import (
	"math/rand"
)

var (
	remainingWeeklyTasks  = []Assignable{}
	remainingMonthlyTasks = []Assignable{}
)

func AssignTasksToAll[T Assignable](tasks []T, members []*Member) {
	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	shuffledMembers := make([]*Member, len(members))
	copy(shuffledMembers, members)
	rand.Shuffle(len(shuffledMembers), func(i, j int) {
		shuffledMembers[i], shuffledMembers[j] = shuffledMembers[j], shuffledMembers[i]
	})

	currentMemberIndex := 0
	for _, task := range tasks {
		task.SetAssignee(shuffledMembers[currentMemberIndex])
		currentMemberIndex++
		if currentMemberIndex >= len(shuffledMembers) {
			currentMemberIndex = 0
		}
	}
}

func AssignTasks[T Assignable](tasks []T, member *Member, amount int) {
	if len(remainingWeeklyTasks) == 0 {
		for _, task := range tasks {
			remainingWeeklyTasks = append(remainingWeeklyTasks, task)
		}
	}

	amountAdded := 0
	for amountAdded < amount && len(remainingWeeklyTasks) > 0 {
		task := remainingWeeklyTasks[0]
		task.SetAssignee(member)
		remainingWeeklyTasks = remainingWeeklyTasks[1:]
		amountAdded++
	}
}

func ClearTasks[T Assignable](tasks []T) {
	for _, task := range tasks {
		task.SetAssignee(nil)
	}
}

func GetAssignedTasks(config *Config, member *Member) []Assignable {
	assignedTasks := []Assignable{}

	for _, task := range config.DailyTasks {
		assignee := task.GetAssignee()
		if assignee != nil && assignee.Name == member.Name {
			assignedTasks = append(assignedTasks, task)
		}
	}

	for _, task := range config.WeeklyTasks {
		assignee := task.GetAssignee()
		if assignee != nil && assignee.Name == member.Name {
			assignedTasks = append(assignedTasks, task)
		}
	}

	for _, task := range config.MonthlyTasks {
		assignee := task.GetAssignee()
		if assignee != nil && assignee.Name == member.Name {
			assignedTasks = append(assignedTasks, task)
		}
	}

	return assignedTasks
}
