package models

import "time"

type Counter struct {
	ID        int64
	Title     string
	Notes     string
	StartDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64
}
