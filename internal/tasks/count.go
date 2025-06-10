// Package tasks presents tasks for the project
package tasks

import (
	"time"

	"github.com/artnikel/iotask/internal/constants"
	"github.com/artnikel/iotask/internal/models"
)

// CountTo200 simulates a long time function
func CountTo200(task *models.Task) {
	task.Status = constants.StatusInProgress
	task.StartedAt = time.Now()

	for i := 1; i <= 200; i++ {
		time.Sleep(1 * time.Second)
		task.CurrentNum = i
	}

	task.Status = constants.StatusCompleted
	task.CompletedAt = time.Now()
}
