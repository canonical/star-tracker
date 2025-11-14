package main

import (
	"encoding/json"
	"net/http"
)

func ProjectsHandler(w http.ResponseWriter, r *http.Request, storage *Storage) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if project.Name == "" || project.UpstreamRepository == "" || project.MirroredRepository == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	storage.AddProject(project)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Project registered successfully",
	})
}
