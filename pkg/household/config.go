package household

import (
	"fmt"

	"github.com/bigkevmcd/go-configparser"
)

type HouseholdConfig struct {
	Members     []*Member
	Tasks       []WeeklyTask
	Assignments map[string][]string
}

func NewHouseholdConfig() *HouseholdConfig {
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

	return &HouseholdConfig{
		Members:     members,
		Tasks:       []WeeklyTask{},
		Assignments: make(map[string][]string),
	}
}
