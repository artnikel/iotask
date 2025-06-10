// Package service contains business logic of a project
package service

import (
	"time"

	"github.com/artnikel/iotask/internal/constants"
	"github.com/artnikel/iotask/internal/models"
	"github.com/artnikel/iotask/internal/tasks"
	"github.com/google/uuid"
)

// TaskService contains Manager struct
type TaskService struct {
	Manager *models.Manager
}

// NewTaskService accepts Manager and returns an object of type *TaskService
func NewTaskService(manager *models.Manager) *TaskService {
	return &TaskService{Manager: manager}
}

// CreateTask creates new Task and insert them into manager
func (s *TaskService) CreateTask() *models.Task {
	s.Manager.Mu.Lock()
	defer s.Manager.Mu.Unlock()

	id := uuid.New()
	task := &models.Task{
		ID:        id,
		Status:    constants.StatusCreated,
		CreatedAt: time.Now(),
	}

	s.Manager.Task[id.String()] = task

	go func(t *models.Task) {
		tasks.CountTo200(t)
	}(task)

	return task
}
