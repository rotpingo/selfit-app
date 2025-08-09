package dto

import (
	"selfit/models"
	"time"
)

type EndTaskDTO struct {
	Notes string `json:"notes"`
}

type TaskResponseDTO struct {
	ID       int64     `json:"id" binding:"required"`
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	IsRepeat bool      `json:"isRepeat" binding:"required"`
	Interval uint      `json:"interval"`
	Notes    string    `json:"notes"`
	DueDate  time.Time `json:"dueDate" binding:"required"`
}

func TaskToResponseDTO(task *models.Task) TaskResponseDTO {
	return TaskResponseDTO{
		ID:       task.ID,
		Title:    task.Title,
		Content:  task.Content,
		IsRepeat: task.IsRepeat,
		Interval: task.Interval,
		Notes:    task.Notes,
	}
}
