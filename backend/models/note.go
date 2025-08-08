package models

import "time"

type Note struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    int64     `json:"userId" binding:"required"`
}

func NewNote(id int64, title string, content string, createdAt time.Time, updatedAt time.Time, userID int64) *Note {
	return &Note{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: updatedAt,
		UserID:    userID,
	}
}
