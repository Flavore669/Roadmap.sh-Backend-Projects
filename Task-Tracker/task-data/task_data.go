package task_data

import "time"

type TaskStatus int64 // enum

const (
	NotStarted TaskStatus = iota
	InProgress
	Completed
)

var TaskStatusToString = map[TaskStatus]string{
	NotStarted: "NotStarted",
	InProgress: "InProgress",
	Completed:  "Completed",
}

type Task struct {
	ID          int       `json:"ID"`
	TaskStatus  string    `json:"TaskStatus"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	Description string    `json:"Description"`
}
