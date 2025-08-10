// Package householdmember provides functionality to manage household members.
//
// It includes the ability to create a new household member with a name and phone number.
package household

type Member struct {
	Name        string
	PhoneNumber string
}

func NewMember(name string, phoneNumber string) *Member {
	return &Member{
		Name:        name,
		PhoneNumber: phoneNumber,
	}
}

type Assignable interface {
	SetAssignee(member *Member)
}

type DailyTask struct {
	Name     string
	Assignee *Member
}

func NewDailyTask(name string) *DailyTask {
	return &DailyTask{
		Name:     name,
		Assignee: nil,
	}
}

func (task *DailyTask) SetAssignee(member *Member) {
	task.Assignee = member
}

type WeeklyTask struct {
	Name     string
	Assignee *Member
}

func NewWeeklyTask(name string) *WeeklyTask {
	return &WeeklyTask{
		Name:     name,
		Assignee: nil,
	}
}

func (task *WeeklyTask) SetAssignee(member *Member) {
	task.Assignee = member
}

type MonthlyTask struct {
	Name     string
	Assignee *Member
}

func NewMonthlyTask(name string) *MonthlyTask {
	return &MonthlyTask{
		Name:     name,
		Assignee: nil,
	}
}

func (task *MonthlyTask) SetAssignee(member *Member) {
	task.Assignee = member
}
