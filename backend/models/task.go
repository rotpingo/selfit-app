package models

import "time"

type Status string

const (
	StatusDone     Status = "done"
	StatusCanceled Status = "canceled"
	StatusProgress Status = "progress"
)

type Task struct {
	ID        uint
	ParentID  uint
	Title     string
	Content   string
	Status    Status
	IsRepeat  bool
	Interval  uint
	Notes     string
	ExecDate  time.Time
	ExecAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
