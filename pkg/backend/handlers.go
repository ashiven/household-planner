package backend

import (
	"encoding/json"
	"fmt"
	"household-planner/pkg/planner"
	"net/http"
)

var Config *planner.Config

func SetConfig(config *planner.Config) {
	Config = config
}

func handleUpdate[T any](w http.ResponseWriter, r *http.Request, section string, setOption func(config *planner.Config, option T)) {
	var updatedItems []T
	if err := json.NewDecoder(r.Body).Decode(&updatedItems); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Config.File.RemoveSection(section)
	Config.File.AddSection(section)

	for _, item := range updatedItems {
		setOption(Config, item)
	}

	if err := Config.File.SaveWithDelimiter(Config.Filename, ":"); err != nil {
		fmt.Fprintf(w, "Error saving config: %v", err)
		return
	}

	// Reload the config to update the in-memory representation
	Config = planner.LoadConfig()
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.Members)
}

func updateMembers(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Members", func(config *planner.Config, member planner.Member) {
		Config.File.Set("Members", member.Name, member.Phonenumber)
	})
}

func getDailyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.DailyTasks)
}

func updateDailyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Daily Tasks", func(config *planner.Config, task planner.DailyTask) {
		Config.File.Set("Daily Tasks", task.Name, "")
	})
}

func getWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.WeeklyTasks)
}

func updateWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Weekly Tasks", func(config *planner.Config, task planner.WeeklyTask) {
		Config.File.Set("Weekly Tasks", task.Name, "")
	})
}

func getMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.MonthlyTasks)
}

func updateMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	handleUpdate(w, r, "Monthly Tasks", func(config *planner.Config, task planner.MonthlyTask) {
		Config.File.Set("Monthly Tasks", task.Name, "")
	})
}
