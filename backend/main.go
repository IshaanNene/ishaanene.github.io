package main

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

func main() {
	http.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		projects := []Project{
			{"GoShell", "A powerful Go-based shell interface.", "Go"},
			{"NetGuard", "Networking toolkit with enhanced security.", "Go"},
			{"Colossus", "ARM emulator with concurrency and memory safety.", "Go"},
			{"VectorX", "High-performance 2D physics engine.", "C"},
			{"sPyC", "Blending Python simplicity with C power.", "C"},
			{"YADTQ", "Distributed task queue using Kafka.", "Python"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(projects)
	})

	http.ListenAndServe(":8080", nil)
}

