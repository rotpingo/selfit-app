package models

import "time"

type Note struct {
	ID        uint
	Title     string    `binding:"required"`
	Content   string    `binding:"required"`
	CreatedAt time.Time `binding:"required"`
	UpdatedAt time.Time `binding:"required"`
	UserID    uint
}

var notes = []Note{}

func (note Note) Save() {
	// add it to database
	notes = append(notes, note)
}

func GetAllNotes() []Note {
	return notes
}
