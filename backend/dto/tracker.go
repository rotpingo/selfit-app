package dto

import (
	"selfit/models"
	"time"
)

type TrackerResponseDTO struct {
	ID            int64     `json:"id" binding:"required"`
	Title         string    `json:"title" binding:"required"`
	Notes         string    `json:"notes"`
	StartDate     time.Time `json:"startDate" binding:"required"`
	CurrentStreak int       `json:"currentStreak"`
	BestStreak    int       `json:"bestStreak"`
}

type CreateTrackerDTO struct {
	Title     string    `json:"title" binding:"required"`
	Notes     string    `json:"notes"`
	StartDate time.Time `json:"startDate"`
}

type UpdateTrackerDTO struct {
	ID    int64  `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Notes string `json:"notes"`
}

func (dto CreateTrackerDTO) ToTrackerModel(userId int64) *models.Tracker {
	return &models.Tracker{
		Title:      dto.Title,
		Notes:      dto.Notes,
		StartDate:  dto.StartDate,
		BestStreak: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		UserID:     userId,
	}
}

func (dto UpdateTrackerDTO) ToTrackerModel(userId int64) *models.Tracker {
	return &models.Tracker{
		ID:        dto.ID,
		Title:     dto.Title,
		Notes:     dto.Notes,
		UpdatedAt: time.Now(),
		UserID:    userId,
	}
}

func TrackerToResponseDTO(tracker *models.Tracker) TrackerResponseDTO {
	return TrackerResponseDTO{
		ID:            tracker.ID,
		Title:         tracker.Title,
		Notes:         tracker.Notes,
		StartDate:     tracker.StartDate,
		CurrentStreak: tracker.CalculateStreak(),
		BestStreak:    tracker.BestStreak,
	}
}
