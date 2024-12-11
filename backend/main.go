package main

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	URL         string `json:"url"`
}

func main() {
	http.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		projects := []Project{
			{"GoShell", "A Go-based CLI with Unix commands.", "Go", "https://github.com/username/GoShell"},
			{"NetGuard", "A secure networking toolkit.", "Go", "https://github.com/username/NetGuard"},
			{"Colossus", "An ARM emulator in Go.", "Go", "https://github.com/username/Colossus"},
			{"VectorX", "A high-performance 2D physics engine.", "C", "https://github.com/username/VectorX"},
			{"sPyC", "Blends Python simplicity with C.", "C", "https://github.com/username/sPyC"},
			{"YADTQ", "Distributed task queue with Kafka.", "Python", "https://github.com/username/YADTQ"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(projects)
	})

	http.ListenAndServe(":8080", nil)
}
