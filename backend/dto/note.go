package dto

import (
	"selfit/models"
	"time"
)

type NoteResponseDTO struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CreateNoteDTO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateNoteDTO struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
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

func (dto UpdateNoteDTO) ToNoteModel(userId int64) *models.Note {
	return &models.Note{
		ID:        dto.ID,
		Title:     dto.Title,
		Content:   dto.Content,
		UpdatedAt: time.Now(),
		UserID:    userId,
	}
}

func NoteToResponseDTO(note *models.Note) NoteResponseDTO {
	return NoteResponseDTO{
		ID:      note.ID,
		Title:   note.Title,
		Content: note.Content,
	}
}
