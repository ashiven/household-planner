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
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/members", getMembers).Methods("GET", "OPTIONS")
	router.HandleFunc("/members", updateMembers).Methods("POST", "OPTIONS")

	router.HandleFunc("/tasks/daily", getDailyTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/tasks/daily", updateDailyTasks).Methods("POST", "OPTIONS")

	router.HandleFunc("/tasks/weekly", getWeeklyTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/tasks/weekly", updateWeeklyTasks).Methods("POST", "OPTIONS")

	router.HandleFunc("/tasks/monthly", getMonthlyTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/tasks/monthly", updateMonthlyTasks).Methods("POST", "OPTIONS")

	router.HandleFunc("/auth", checkAdminPassword).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", router))
}
