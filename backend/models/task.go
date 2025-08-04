package models

import "time"

type Status string

const (
	StatusDone     Status = "done"
	StatusAborted  Status = "aborted"
	StatusProgress Status = "progress"
)

type Task struct {
	ID        int64     `json:"id"`
	ParentID  int64     `json:"parent_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    Status    `json:"status"`
	IsRepeat  bool      `json:"isRepeat"`
	Interval  uint      `json:"interval"`
	Notes     string    `json:"notes"`
	DueDate   time.Time `json:"due_date"`
	ExecAt    time.Time `json:"exec_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
