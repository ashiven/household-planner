// Package backend implements the server for the household planner application.
package backend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	fmt.Println("[INFO] Starting Household Planner API server...")
	router := mux.NewRouter()
	router.HandleFunc("/members", getMembers).Methods("GET")
	router.HandleFunc("/members", updateMembers).Methods("POST")

	router.HandleFunc("/tasks/daily", getDailyTasks).Methods("GET")
	router.HandleFunc("/tasks/daily", updateDailyTasks).Methods("POST")

	router.HandleFunc("/tasks/weekly", getWeeklyTasks).Methods("GET")
	router.HandleFunc("/tasks/weekly", updateWeeklyTasks).Methods("POST")

	router.HandleFunc("/tasks/monthly", getMonthlyTasks).Methods("GET")
	router.HandleFunc("/tasks/monthly", updateMonthlyTasks).Methods("POST")

	router.HandleFunc("/auth", checkAdminPassword).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
