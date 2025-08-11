package backend

import (
	"encoding/json"
	"household-planner/pkg/planner"
	"net/http"
)

func getMembers(w http.ResponseWriter, r *http.Request) {
	config := planner.LoadConfig()
	json.NewEncoder(w).Encode(config.Members)
}

func updateMembers(w http.ResponseWriter, r *http.Request) {
	updatedMembers := []planner.Member{}
	err := json.NewDecoder(r.Body).Decode(&updatedMembers)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	config := planner.LoadConfig()
	config.File.RemoveSection("Members")
	config.File.AddSection("Members")
	for _, member := range updatedMembers {
		config.File.Set("Members", member.Name, member.Phonenumber)
	}
	config.File.SaveWithDelimiter(config.Filename, ":")
}

func getDailyTasks(w http.ResponseWriter, r *http.Request) {
	config := planner.LoadConfig()
	json.NewEncoder(w).Encode(config.DailyTasks)
}

func updateDailyTasks(w http.ResponseWriter, r *http.Request) {
	updatedDailyTasks := []planner.DailyTask{}
	err := json.NewDecoder(r.Body).Decode(&updatedDailyTasks)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	config := planner.LoadConfig()
	config.File.RemoveSection("Daily Tasks")
	config.File.AddSection("Daily Tasks")
	for _, dailyTask := range updatedDailyTasks {
		config.File.Set("Daily Tasks", dailyTask.Name, "")
	}
	config.File.SaveWithDelimiter(config.Filename, ":")
}

func getWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	config := planner.LoadConfig()
	json.NewEncoder(w).Encode(config.WeeklyTasks)
}

func updateWeeklyTasks(w http.ResponseWriter, r *http.Request) {
	updatedWeeklyTasks := []planner.WeeklyTask{}
	err := json.NewDecoder(r.Body).Decode(&updatedWeeklyTasks)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	config := planner.LoadConfig()
	config.File.RemoveSection("Weekly Tasks")
	config.File.AddSection("Weekly Tasks")
	for _, weeklyTask := range updatedWeeklyTasks {
		config.File.Set("Weekly Tasks", weeklyTask.Name, "")
	}
	config.File.SaveWithDelimiter(config.Filename, ":")
}

func getMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	config := planner.LoadConfig()
	json.NewEncoder(w).Encode(config.MonthlyTasks)
}

func updateMonthlyTasks(w http.ResponseWriter, r *http.Request) {
	updatedMonthlyTasks := []*planner.MonthlyTask{}
	err := json.NewDecoder(r.Body).Decode(&updatedMonthlyTasks)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	config := planner.LoadConfig()
	config.File.RemoveSection("Monthly Tasks")
	config.File.AddSection("Monthly Tasks")
	for _, monthlyTask := range updatedMonthlyTasks {
		config.File.Set("Monthly Tasks", monthlyTask.Name, "")
	}
	config.File.SaveWithDelimiter(config.Filename, ":")
}
