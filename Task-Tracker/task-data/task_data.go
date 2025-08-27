package task_data

import (
	"errors"
	"time"
)

var validStatus = []string{"not-started", "in-progress", "done"}

func IsValidStatus(status string) error {
	for i := range validStatus {
		if validStatus[i] == status {
			return nil
		}
	}
	return errors.New("not a valid status")
}

type Task struct {
	ID          int       `json:"ID"`
	TaskStatus  string    `json:"TaskStatus"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	Description string    `json:"Description"`
}
