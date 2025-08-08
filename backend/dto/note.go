package dto

import (
	"selfit/models"
	"time"
)

type CreateNoteDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (dto CreateNoteDTO) ToNoteModel(userId int64) *models.Note {
	return &models.Note{
		Title:     dto.Title,
		Content:   dto.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userId,
	}
}
