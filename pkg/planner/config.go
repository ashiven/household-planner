package planner

import (
	"fmt"

	"github.com/bigkevmcd/go-configparser"
)

const configPath = "config.ini"

type Config struct {
	Filename     string
	File         *configparser.ConfigParser
	Members      []*Member
	DailyTasks   []*DailyTask
	WeeklyTasks  []*WeeklyTask
	MonthlyTasks []*MonthlyTask
}

func LoadConfig() *Config {
	parser, err := configparser.NewConfigParserFromFile(configPath)
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
		Filename:     configPath,
		File:         parser,
		Members:      members,
		DailyTasks:   dailyTasks,
		WeeklyTasks:  weeklyTasks,
		MonthlyTasks: monthlyTasks,
	}
}
