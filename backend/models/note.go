package models

import "time"

type Note struct {
	ID        int64
	Title     string    `binding:"required"`
	Content   string    `binding:"required"`
	CreatedAt time.Time `binding:"required"`
	UpdatedAt time.Time `binding:"required"`
	UserID    uint
}
