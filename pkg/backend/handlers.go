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

func getMembers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.Members)
}

func updateMembers(w http.ResponseWriter, r *http.Request) {
	updatedMembers := []planner.Member{}
	err := json.NewDecoder(r.Body).Decode(&updatedMembers)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Config.File.RemoveSection("Members")
	Config.File.AddSection("Members")
	for _, member := range updatedMembers {
		Config.File.Set("Members", member.Name, member.Phonenumber)
	}
	Config.File.SaveWithDelimiter(Config.Filename, ":")

	// Reload the config to update the in-memory representation
	Config = planner.LoadConfig()
}

func getDailyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.DailyTasks)
}

func updateDailyTasks(w http.ResponseWriter, r *http.Request) {
	updatedDailyTasks := []planner.DailyTask{}
	err := json.NewDecoder(r.Body).Decode(&updatedDailyTasks)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Config.File.RemoveSection("Daily Tasks")
	Config.File.AddSection("Daily Tasks")
	for _, dailyTask := range updatedDailyTasks {
		Config.File.Set("Daily Tasks", dailyTask.Name, "")
	}
	Config.File.SaveWithDelimiter(Config.Filename, ":")
	Config = planner.LoadConfig()
}

func getWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.WeeklyTasks)
}

func updateWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	updatedWeeklyTasks := []planner.WeeklyTask{}
	err := json.NewDecoder(r.Body).Decode(&updatedWeeklyTasks)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Config.File.RemoveSection("Weekly Tasks")
	Config.File.AddSection("Weekly Tasks")
	for _, weeklyTask := range updatedWeeklyTasks {
		Config.File.Set("Weekly Tasks", weeklyTask.Name, "")
	}
	Config.File.SaveWithDelimiter(Config.Filename, ":")
	Config = planner.LoadConfig()
}

func getMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Config.MonthlyTasks)
}

func updateMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	updatedMonthlyTasks := []*planner.MonthlyTask{}
	err := json.NewDecoder(r.Body).Decode(&updatedMonthlyTasks)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Config.File.RemoveSection("Monthly Tasks")
	Config.File.AddSection("Monthly Tasks")
	for _, monthlyTask := range updatedMonthlyTasks {
		Config.File.Set("Monthly Tasks", monthlyTask.Name, "")
	}
	Config.File.SaveWithDelimiter(Config.Filename, ":")
	Config = planner.LoadConfig()
}
