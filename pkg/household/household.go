// Package household provides functionality to manage household members.
//
// It includes the ability to create a new household member with a name and phone number.
package household

import "math/rand"

func AssignTasksToAll[T Assignable](tasks []T, members []*Member) {
	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})

	currentMemberIndex := 0
	for _, task := range tasks {
		task.SetAssignee(members[currentMemberIndex])
		currentMemberIndex++
		if currentMemberIndex >= len(members) {
			currentMemberIndex = 0
		}
	}
}

func AssignTasks[T Assignable](tasks []T, member *Member) {
	// TODO: weekly tasks are dynamically split into daily tasks so each member gets a fair share
	// TODO: monthly tasks are split into daily tasks with spaced intervals so they are evenly spread across the month
}

func ClearTasks[T Assignable](tasks []T) {
	for _, task := range tasks {
		task.SetAssignee(nil)
	}
}

func GetAssignedTasks[T Assignable](tasks []T, member *Member) []T {
	assignedTasks := []T{}
	for _, task := range tasks {
		assignee := task.GetAssignee()
		if assignee != nil && assignee.Name == member.Name {
			assignedTasks = append(assignedTasks, task)
		}
	}
	return assignedTasks
}
