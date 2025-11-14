package main

import (
	"sync"
)

type Storage struct {
	mu         sync.RWMutex
	projects   []Project
	components []Component
}

func NewStorage() *Storage {
	return &Storage{
		projects:   make([]Project, 0),
		components: make([]Component, 0),
	}
}

func (s *Storage) AddProject(project Project) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.projects = append(s.projects, project)
}

func (s *Storage) AddComponent(component Component) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.components = append(s.components, component)
}

func (s *Storage) SearchComponents(name, version, sourcecraftName, sourcecraftTrack, license string) []Component {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []Component
	for _, component := range s.components {
		if matchesSearch(component, name, version, sourcecraftName, sourcecraftTrack, license) {
			results = append(results, component)
		}
	}
	return results
}
