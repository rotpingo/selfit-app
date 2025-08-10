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

type CreateTaskDTO struct {
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	IsRepeat bool      `json:"isRepeat" binding:"required"`
	Interval uint      `json:"interval"`
	DueDate  time.Time `json:"dueDate" binding:"required"`
}

type UpdateTaskDTO struct {
	ID       int64     `json:"id" binding:"required"`
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	IsRepeat bool      `json:"isRepeat" binding:"required"`
	Interval uint      `json:"interval"`
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

func (dto CreateTaskDTO) ToTaskModel(userId int64) *models.Task {
	return &models.Task{
		Title:     dto.Title,
		Content:   dto.Content,
		Status:    models.StatusProgress,
		IsRepeat:  dto.IsRepeat,
		Interval:  dto.Interval,
		DueDate:   dto.DueDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userId,
	}
}

func (dto UpdateTaskDTO) ToTaskModel(userId int64) *models.Task {
	return &models.Task{
		ID:        dto.ID,
		Title:     dto.Title,
		Content:   dto.Content,
		IsRepeat:  dto.IsRepeat,
		Interval:  dto.Interval,
		DueDate:   dto.DueDate,
		UpdatedAt: time.Now(),
		UserID:    userId,
	}
}
