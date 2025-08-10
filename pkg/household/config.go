package household

import (
	"fmt"
	"math/rand"

	"github.com/bigkevmcd/go-configparser"
)

type Config struct {
	Members      []*Member
	DailyTasks   []*DailyTask
	WeeklyTasks  []*WeeklyTask
	MonthlyTasks []*MonthlyTask
}

func NewConfig() *Config {
	parser, err := configparser.NewConfigParserFromFile("config.ini")
	if err != nil {
		fmt.Println(err.Error())
	}

	memberInfo, err := parser.Items("Members")
	if err != nil {
		fmt.Println(err.Error())
	}
	members := []*Member{}
	for memberName, phoneNumeber := range memberInfo {
		members = append(members, NewMember(memberName, phoneNumeber))
	}

	dailyTaskInfo, err := parser.Options("Daily Tasks")
	if err != nil {
		fmt.Println(err.Error())
	}
	dailyTasks := []*DailyTask{}
	for _, dailyTask := range dailyTaskInfo {
		dailyTasks = append(dailyTasks, NewDailyTask(dailyTask))
	}

	weeklyTaskInfo, err := parser.Options("Weekly Tasks")
	if err != nil {
		fmt.Println(err.Error())
	}
	weeklyTasks := []*WeeklyTask{}
	for _, weeklyTask := range weeklyTaskInfo {
		weeklyTasks = append(weeklyTasks, NewWeeklyTask(weeklyTask))
	}

	monthlyTaskInfo, err := parser.Options("Monthly Tasks")
	if err != nil {
		fmt.Println(err.Error())
	}
	monthlyTasks := []*MonthlyTask{}
	for _, monthlyTask := range monthlyTaskInfo {
		monthlyTasks = append(monthlyTasks, NewMonthlyTask(monthlyTask))
	}

	return &Config{
		Members:      members,
		DailyTasks:   dailyTasks,
		WeeklyTasks:  weeklyTasks,
		MonthlyTasks: monthlyTasks,
	}
}

func AssignTasks[T Assignable](tasks []T, members []*Member) {
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
