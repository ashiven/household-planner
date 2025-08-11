package planner

type Household struct {
	config                *Config
	Members               []*Member
	DailyTasks            []*DailyTask
	WeeklyTasks           []*WeeklyTask
	MonthlyTasks          []*MonthlyTask
	weeklyTasksPerDay     int
	remainingWeeklyTasks  []*WeeklyTask
	monthlyTasksPerDay    int
	remainingMonthlyTasks []*MonthlyTask
	dayOfTheMonth         int
	taskIntervalMonth     int
}

func NewHousehold(config *Config) *Household {
	return &Household{
		config,
		config.Members,
		config.DailyTasks,
		config.WeeklyTasks,
		config.MonthlyTasks,
		min(len(config.WeeklyTasks)/len(config.Members), 1),
		[]*WeeklyTask{},
		min(len(config.MonthlyTasks)/len(config.Members), 1),
		[]*MonthlyTask{},
		1,
		30 / len(config.MonthlyTasks),
	}
}

type HouseholdInterface interface {
	AssignDailyTasks(household *Household)
	AssignWeeklyTasks(household *Household)
	AssignMonthlyTasks(household *Household)
	GetAssignedTasks(household *Household, member *Member) []Assignable
	ClearAssignments(household *Household)
}

type Member struct {
	Name        string
	Phonenumber string
}

func NewMember(name string, phonenumber string) *Member {
	return &Member{
		Name:        name,
		Phonenumber: phonenumber,
	}
}

type Assignable interface {
	SetAssignee(member *Member)
	GetAssignee() *Member
	GetName() string
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

func (task *DailyTask) GetAssignee() *Member {
	return task.Assignee
}

func (task *DailyTask) GetName() string {
	return task.Name
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

func (task *WeeklyTask) GetAssignee() *Member {
	return task.Assignee
}

func (task *WeeklyTask) GetName() string {
	return task.Name
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

func (task *MonthlyTask) GetAssignee() *Member {
	return task.Assignee
}

func (task *MonthlyTask) GetName() string {
	return task.Name
}
