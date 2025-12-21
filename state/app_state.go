package state

import "sync"

type Task struct {
	ID int
	Name string
	Status string
}

type AppState struct {
	mu sync.Mutex
	Tasks []Task
	Logs []string
}

func NewAppState() *AppState {
	return &AppState{
		Tasks: []Task{
			{1, "Initialize", "Done"},
			{2, "Process", "Pending"},
			{3, "Finalize", "Pending"},
		},
	}
}

func (s *AppState) AddLog(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Logs = append(s.Logs, msg)
}

func (s *AppState) UpdateTask(id int, status string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			s.Tasks[i].Status = status
		}
	}
}