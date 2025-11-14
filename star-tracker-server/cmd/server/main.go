package main

import (
	"log"
	"net/http"
)

func main() {
	storage := NewStorage()

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		ProjectsHandler(w, r, storage)
	})
	http.HandleFunc("/components", func(w http.ResponseWriter, r *http.Request) {
		ComponentsHandler(w, r, storage)
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
