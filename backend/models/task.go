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
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Status    Status    `json:"status" binding:"required"`
	IsRepeat  bool      `json:"isRepeat" binding:"required"`
	Interval  uint      `json:"interval"`
	Notes     string    `json:"notes"`
	DueDate   time.Time `json:"dueDate" binding:"required"`
	ExecAt    time.Time `json:"execAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    int64     `json:"userId" binding:"required"`
}
