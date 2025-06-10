package tasks

import (
	"testing"
	"time"

	"github.com/artnikel/iotask/internal/constants"
	"github.com/artnikel/iotask/internal/models"
)

func TestCountTo100(t *testing.T) {
	task := &models.Task{}

	go CountTo200(task)

	time.Sleep(2 * time.Second)

	if task.Status != constants.StatusInProgress && task.Status != constants.StatusCompleted {
		t.Errorf("unexpected task status: %s", task.Status)
	}
}
