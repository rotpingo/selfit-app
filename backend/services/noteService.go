package services

import (
	"selfit/database"
	"selfit/models"
)

func GetAllNotes() ([]models.Note, error) {
	query := "SELECT * FROM notes"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.UserID)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func SaveNote(note models.Note) error {
	query := `
	INSERT INTO notes(title, content) 
	VALUES (?, ?)
	`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(note.Title, note.Content)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}
	note.ID = id

	return err
}
