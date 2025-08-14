// Package planner provides functionality to manage household members.
//
// It includes the ability to create a new household member with a name and phone number.
package planner

import (
	"errors"

	"github.com/bigkevmcd/go-configparser"
)

const configPath = "config.ini"

type Member struct {
	Name        string
	Phonenumber string
}

type Household struct {
	Configfile            string
	Config                *configparser.ConfigParser
	Members               []*Member
	DailyTasks            []*DailyTask
	WeeklyTasks           []*WeeklyTask
	MonthlyTasks          []*MonthlyTask
	dayOfTheMonth         int
	currentMemberIndex    int
	currentMember         *Member
	remainingWeeklyTasks  []*WeeklyTask
	remainingMonthlyTasks []*MonthlyTask
}

func NewMember(name string, phonenumber string) *Member {
	return &Member{
		Name:        name,
		Phonenumber: phonenumber,
	}
}

func NewHousehold() (*Household, error) {
	parser, err := configparser.NewConfigParserFromFile(configPath)
	if err != nil {
		return nil, err
	}

	memberInfo, err := parser.Items("Members")
	if err != nil {
		return nil, err
	}
	members := []*Member{}
	for memberName, phoneNumeber := range memberInfo {
		members = append(members, NewMember(memberName, phoneNumeber))
	}

	dailyTaskInfo, err := parser.Options("Daily Tasks")
	if err != nil {
		return nil, err
	}
	dailyTasks := []*DailyTask{}
	for _, dailyTask := range dailyTaskInfo {
		dailyTasks = append(dailyTasks, NewDailyTask(dailyTask))
	}

	weeklyTaskInfo, err := parser.Options("Weekly Tasks")
	if err != nil {
		return nil, err
	}
	weeklyTasks := []*WeeklyTask{}
	for _, weeklyTask := range weeklyTaskInfo {
		weeklyTasks = append(weeklyTasks, NewWeeklyTask(weeklyTask))
	}

	monthlyTaskInfo, err := parser.Options("Monthly Tasks")
	if err != nil {
		return nil, err
	}
	monthlyTasks := []*MonthlyTask{}
	for _, monthlyTask := range monthlyTaskInfo {
		monthlyTasks = append(monthlyTasks, NewMonthlyTask(monthlyTask))
	}

	if len(members) == 0 || len(dailyTasks) == 0 || len(weeklyTasks) == 0 || len(monthlyTasks) == 0 {
		return nil, errors.New("household must have at least one member and one task in each category")
	}

	return &Household{
		Configfile:   configPath,
		Config:       parser,
		Members:      members,
		DailyTasks:   dailyTasks,
		WeeklyTasks:  weeklyTasks,
		MonthlyTasks: monthlyTasks,
	}, nil
}
