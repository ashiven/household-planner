package backend

import (
	"encoding/json"
	"household-planner/pkg/planner"
	"net/http"
	"sync"
)

var (
	Config   *planner.Config
	FileLock sync.Mutex
)

func SetConfig(config *planner.Config) {
	Config = config
}

func handleUpdate[T any](w http.ResponseWriter, r *http.Request, section string, setConfigOption func(option T), setOptionsMemory func(updated []*T)) {
	FileLock.Lock()
	defer FileLock.Unlock()

	var updatedOptions []T
	if err := json.NewDecoder(r.Body).Decode(&updatedOptions); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// We need to make sure that not only the config file is updated, but also the in-memory representation.
	// We do this by creating an in-memory representation of the updated options and returning them to the caller who defines
	// how this should be used to update the in-memory state. (via setOptionsMemory)
	configOptionsMemory := make([]*T, len(updatedOptions))
	for optionIndex, option := range updatedOptions {
		configOptionsMemory[optionIndex] = &option
	}
	setOptionsMemory(configOptionsMemory)

	Config.File.RemoveSection(section)
	Config.File.AddSection(section)

	for _, option := range updatedOptions {
		setConfigOption(option)
	}

	if err := Config.File.SaveWithDelimiter(Config.Filename, ":"); err != nil {
		http.Error(w, "Error saving config: %v", http.StatusInternalServerError)
		return
	}
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.Members)
}

func updateMembers(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Members", func(updatedMember planner.Member) {
		Config.File.Set("Members", updatedMember.Name, updatedMember.Phonenumber)
	}, func(updatedMembers []*planner.Member) {
		Config.Members = updatedMembers
	})
}

func getDailyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.DailyTasks)
}

func updateDailyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Daily Tasks", func(task planner.DailyTask) {
		Config.File.Set("Daily Tasks", task.Name, "")
	}, func(updatedTasks []*planner.DailyTask) {
		Config.DailyTasks = updatedTasks
	})
}

func getWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.WeeklyTasks)
}

func updateWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Weekly Tasks", func(task planner.WeeklyTask) {
		Config.File.Set("Weekly Tasks", task.Name, "")
	}, func(updatedTasks []*planner.WeeklyTask) {
		Config.WeeklyTasks = updatedTasks
	})
}

func getMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.MonthlyTasks)
}

func updateMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Monthly Tasks", func(task planner.MonthlyTask) {
		Config.File.Set("Monthly Tasks", task.Name, "")
	}, func(updatedTasks []*planner.MonthlyTask) {
		Config.MonthlyTasks = updatedTasks
	})
}
