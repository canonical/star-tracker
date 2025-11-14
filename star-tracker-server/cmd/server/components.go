package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func ComponentsHandler(w http.ResponseWriter, r *http.Request, storage *Storage) {
	switch r.Method {
	case http.MethodPost:
		registerComponent(w, r, storage)
	case http.MethodGet:
		searchComponents(w, r, storage)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func registerComponent(w http.ResponseWriter, r *http.Request, storage *Storage) {
	var component Component
	if err := json.NewDecoder(r.Body).Decode(&component); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if component.Name == "" || component.Version == "" || component.Project == "" ||
		component.SourcecraftName == "" || component.SourcecraftTrack == "" || component.License == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	storage.AddComponent(component)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Component registered successfully",
	})
}

func searchComponents(w http.ResponseWriter, r *http.Request, storage *Storage) {
	query := r.URL.Query()
	name := query.Get("name")
	version := query.Get("version")
	sourcecraftName := query.Get("sourcecraft_name")
	sourcecraftTrack := query.Get("sourcecraft_track")
	license := query.Get("license")

	results := storage.SearchComponents(name, version, sourcecraftName, sourcecraftTrack, license)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":   len(results),
		"results": results,
	})
}

func matchesSearch(component Component, name, version, sourcecraftName, sourcecraftTrack, license string) bool {
	if name != "" && !strings.Contains(strings.ToLower(component.Name), strings.ToLower(name)) {
		return false
	}
	if version != "" && !strings.Contains(strings.ToLower(component.Version), strings.ToLower(version)) {
		return false
	}
	if sourcecraftName != "" && !strings.Contains(strings.ToLower(component.SourcecraftName), strings.ToLower(sourcecraftName)) {
		return false
	}
	if sourcecraftTrack != "" && !strings.Contains(strings.ToLower(component.SourcecraftTrack), strings.ToLower(sourcecraftTrack)) {
		return false
	}
	if license != "" && !strings.Contains(strings.ToLower(component.License), strings.ToLower(license)) {
		return false
	}
	return true
}
