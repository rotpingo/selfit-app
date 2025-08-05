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
	ParentID  int64     `json:"parentId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    Status    `json:"status"`
	IsRepeat  bool      `json:"isRepeat"`
	Interval  uint      `json:"interval"`
	Notes     string    `json:"notes"`
	DueDate   time.Time `json:"dueDate"`
	ExecAt    time.Time `json:"execAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    int64     `json:"userId"`
}
