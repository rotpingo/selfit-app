package services

import (
	"fmt"
	"selfit/database"
	"selfit/models"
	"time"
)

func GetAllNotes() ([]models.Note, error) {
	query := "SELECT * FROM notes"
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("error fetching:", err)
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.UserID)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func CreateNote(note *models.Note) error {
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	query := `
	INSERT INTO notes(title, content, created_at, updated_at, user_id) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	err := database.DB.QueryRow(
		query,
		note.Title,
		note.Content,
		note.CreatedAt,
		note.UpdatedAt,
		note.UserID,
	).Scan(&note.ID)

	if err != nil {
		fmt.Println("insert error:", err)
		return err
	}

	return nil
}

func UpdateNote(note *models.Note) error {
	note.UpdatedAt = time.Now()

	query := `
		UPDATE notes
		SET title = $1, content = $2, updated_at = $3
		WHERE id = $4
	`

	_, err := database.DB.Exec(query, note.Title, note.Content, note.UpdatedAt, note.ID)
	if err != nil {
		fmt.Println("update error:", err)
		return err
	}

	return nil
}

func DeleteNoteById(id int) error {

	query := `DELETE FROM notes WHERE id = $1`

	_, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete note: %w", err)
	}

	return nil
}
