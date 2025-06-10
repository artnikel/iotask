// Package models provides the data models used in the application
package models

import (
	"sync"
	"time"

	"github.com/artnikel/iotask/internal/constants"
	"github.com/google/uuid"
)

// Task entity
type Task struct {
	ID          uuid.UUID
	Status      constants.TaskStatus
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
	CurrentNum  int
}

// Manager entity
type Manager struct {
	Task map[string]*Task
	Mu   sync.Mutex
}
