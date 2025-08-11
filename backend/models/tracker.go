package models

import (
	"time"
)

type Tracker struct {
	ID            int64     `json:"id" binding:"required"`
	Title         string    `json:"title" binding:"required"`
	Notes         string    `json:"notes"`
	StartDate     time.Time `json:"startDate" binding:"required"`
	CurrentStreak int       `json:"currentStrak"`
	BestStreak    int       `json:"bestStreak"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"uptadedAt"`
	UserID        int64     `json:"userId"`
}

func (tracker *Tracker) CalculateStreak() int {
	now := time.Now()
	diff := now.Sub(tracker.StartDate)
	return int(diff.Hours() / 24)

}
