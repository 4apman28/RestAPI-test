package models

import "time"

type Status string

const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	CompletedAt *time.Time
	Deadline    *time.Time
}
