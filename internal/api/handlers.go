// Package api contains HTTP handlers for task management
package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/artnikel/iotask/internal/service"
)

// Handler provides HTTP endpoints backed by a Service
type Handler struct {
	Service *service.TaskService
}

// NewHandler creates a new Handler with the given service object
func NewHandler(s *service.TaskService) *Handler {
	return &Handler{Service: s}
}

// CreateTaskHandler handles POST requests to add a new ping task
func (h *Handler) CreateTaskHandler(w http.ResponseWriter, _ *http.Request) {
	task := h.Service.CreateTask()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"task_id": task.ID.String()})
}

// GetTaskHandler handles GET requests to retrieve task status by ID
func (h *Handler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")

	h.Service.Manager.Mu.Lock()
	task, exists := h.Service.Manager.Task[id]
	h.Service.Manager.Mu.Unlock()

	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	status := map[string]interface{}{
		"task_id":          task.ID,
		"status":           task.Status,
		"created_at":       task.CreatedAt.Format(time.RFC3339),
		"current_num":      task.CurrentNum,
		"duration_seconds": time.Since(task.CreatedAt).Seconds(),
	}

	if !task.CompletedAt.IsZero() {
		status["completed_at"] = task.CompletedAt.Format(time.RFC3339)
		status["total_duration_seconds"] = task.CompletedAt.Sub(task.CreatedAt).Seconds()
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(status)
}

// DeleteTaskHandler handles DELETE requests to retrieve task by ID
func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")

	h.Service.Manager.Mu.Lock()
	defer h.Service.Manager.Mu.Unlock()

	if _, ok := h.Service.Manager.Task[id]; !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	delete(h.Service.Manager.Task, id)
	w.WriteHeader(http.StatusNoContent)
}
