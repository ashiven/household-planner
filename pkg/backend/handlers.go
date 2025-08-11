package backend

import (
	"encoding/json"
	"household-planner/pkg/planner"
	"net/http"
)

var Config *planner.Config

func SetConfig(config *planner.Config) {
	Config = config
}

func handleUpdate[T any](w http.ResponseWriter, r *http.Request, section string, setOption func(option T)) {
	var updatedOptions []T
	if err := json.NewDecoder(r.Body).Decode(&updatedOptions); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Config.File.RemoveSection(section)
	Config.File.AddSection(section)

	for _, option := range updatedOptions {
		setOption(option)
	}

	if err := Config.File.SaveWithDelimiter(Config.Filename, ":"); err != nil {
		http.Error(w, "Error saving config: %v", http.StatusInternalServerError)
		return
	}

	// Reload the config to update the in-memory representation
	Config = planner.LoadConfig()
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.Members)
}

func updateMembers(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Members", func(member planner.Member) {
		Config.File.Set("Members", member.Name, member.Phonenumber)
	})
}

func getDailyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.DailyTasks)
}

func updateDailyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Daily Tasks", func(task planner.DailyTask) {
		Config.File.Set("Daily Tasks", task.Name, "")
	})
}

func getWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.WeeklyTasks)
}

func updateWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Weekly Tasks", func(task planner.WeeklyTask) {
		Config.File.Set("Weekly Tasks", task.Name, "")
	})
}

func getMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.MonthlyTasks)
}

func updateMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Monthly Tasks", func(task planner.MonthlyTask) {
		Config.File.Set("Monthly Tasks", task.Name, "")
	})
}
