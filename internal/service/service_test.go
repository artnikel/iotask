package service

import (
	"testing"

	"github.com/artnikel/iotask/internal/models"
)

func TestCreateTask(t *testing.T) {
	manager := &models.Manager{
		Task: make(map[string]*models.Task),
	}
	s := NewTaskService(manager)
	task := s.CreateTask()

	if task == nil {
		t.Fatal("expected task, got nil")
	}
	if task.Status != "created" && task.Status != "in_progress" {
		t.Errorf("unexpected status: %s", task.Status)
	}
	if task.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}
}
