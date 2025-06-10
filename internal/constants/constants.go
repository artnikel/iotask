// Package constants defines shared constants used across the application
package constants

import "time"

// TaskStatus represents the state of a task
type TaskStatus string

const (
	// StatusCreated - Task is created
	StatusCreated TaskStatus = "created"
	// StatusInProgress - Task is in progress
	StatusInProgress TaskStatus = "in_progress"
	// StatusCompleted - Task completed successfully
	StatusCompleted TaskStatus = "completed"
	// StatusFailed - Task execution failed
	StatusFailed TaskStatus = "failed"
	// ServerTimeout is read and write timeout of server config
	ServerTimeout = 10 * time.Second
	// DirPerm - Directory permission
	DirPerm = 0o750
	// FilePerm - File permission
	FilePerm = 0o600
)
