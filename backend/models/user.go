package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64          `json:"id"`
	Name      sql.NullString `json:"name"`
	Email     string         `json:"email" binding:"required"`
	Password  string         `json:"password" binding:"required"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
