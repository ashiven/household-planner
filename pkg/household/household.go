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

type WeeklyTask struct {
	Name string
	Frequency int
}

func NewWeeklyTask(name string, frequency int) *WeeklyTask {
return &WeeklyTask{
		Name: name,
		Frequency: frequency,
	}
}
