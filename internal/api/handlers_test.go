package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/artnikel/iotask/internal/models"
	"github.com/artnikel/iotask/internal/service"
)

func setupTestHandler() *Handler {
	manager := &models.Manager{
		Task: make(map[string]*models.Task),
	}
	svc := service.NewTaskService(manager)
	return NewHandler(svc)
}

func TestCreateTaskHandler(t *testing.T) {
	h := setupTestHandler()

	req := httptest.NewRequest(http.MethodPost, "/tasks", http.NoBody)
	w := httptest.NewRecorder()

	h.CreateTaskHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}

	var body map[string]string
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		t.Fatal("failed to decode response:", err)
	}

	if _, ok := body["task_id"]; !ok {
		t.Error("response missing task_id")
	}
}

func TestGetTaskHandler_NotFound(t *testing.T) {
	h := setupTestHandler()

	req := httptest.NewRequest(http.MethodGet, "/tasks/invalid-id", http.NoBody)
	w := httptest.NewRecorder()

	h.GetTaskHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestDeleteTaskHandler(t *testing.T) {
	h := setupTestHandler()

	req := httptest.NewRequest(http.MethodPost, "/tasks", http.NoBody)
	w := httptest.NewRecorder()
	h.CreateTaskHandler(w, req)

	var body map[string]string
	json.NewDecoder(w.Body).Decode(&body)
	taskID := body["task_id"]

	delReq := httptest.NewRequest(http.MethodDelete, "/tasks/"+taskID, http.NoBody)
	delW := httptest.NewRecorder()
	h.DeleteTaskHandler(delW, delReq)

	if delW.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", delW.Code)
	}
}
