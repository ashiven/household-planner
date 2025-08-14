// Package backend implements the server for the household planner application.
package backend

import (
	"fmt"
	"household-planner/pkg/planner"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	domainName    = planner.GetEnvVar("DOMAIN_NAME")
	allowedOrigin = fmt.Sprintf("https://%s", domainName)
)

func StartServer() {
	fmt.Println("[INFO] Starting Household Planner API server...")

	router := mux.NewRouter()

	if domainName == "" {
		allowedOrigin = "http://localhost"
	}

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
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
