package services

import (
	"time"
)

// Task represents a single task with its status
type Task struct {
	ID     int
	Title  string
	Status string // e.g., "completed", "in-progress", "pending"
	DueDate time.Time
}

// TaskService handles operations related to tasks
type TaskService struct {
	Tasks []Task
}

// GetStatistics returns statistics about tasks
func (s *TaskService) GetStatistics() map[string]int {
	completed := 0
	inProgress := 0
	pending := 0

	for _, task := range s.Tasks {
		switch task.Status {
		case "completed":
			completed++
		case "in-progress":
			inProgress++
		case "pending":
			pending++
		}
	}

	return map[string]int{
		"completed":   completed,
		"in-progress": inProgress,
		"pending":     pending,
	}
}

// GetTasksByStatus returns tasks filtered by their status
func (s *TaskService) GetTasksByStatus(status string) []Task {
	var filteredTasks []Task
	for _, task := range s.Tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

// CalculateCompletion returns the completion percentage of tasks
func (s *TaskService) CalculateCompletion() float64 {
	if len(s.Tasks) == 0 {
		return 0.0
	}
	completedTasks := s.GetStatistics()["completed"]
	return float64(completedTasks) / float64(len(s.Tasks)) * 100
}

// GenerateTaskReport generates a report of tasks
func (s *TaskService) GenerateTaskReport() string {
	report := "Task Report:\n"
	for _, task := range s.Tasks {
		report += task.Title + " - Status: " + task.Status + "\n"
	}
	return report
}