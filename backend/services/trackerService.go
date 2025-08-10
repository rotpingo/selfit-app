package services

import (
	"fmt"
	"selfit/database"
	"selfit/models"
)

func CreateTracker(tracker *models.Tracker) error {

	query := `
	INSERT INTO tracker(title, notes, start_date, best_streak, created_at, updated_at, user_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
	`
	err := database.DB.QueryRow(
		query,
		tracker.Title,
		tracker.Notes,
		tracker.StartDate,
		0,
		tracker.CreatedAt,
		tracker.UpdatedAt,
		tracker.UserID,
	).Scan(&tracker.ID)

	if err != nil {
		fmt.Println("insert error:", err)
		return err
	}

	return nil
}

func UpdateTracker(tracker *models.Tracker) error {
	query := `
		UPDATE tracker
		SET title = $1, notes = $2, updated_at = $3
		WHERE id = $4
		AND user_id = $5
	`

	_, err := database.DB.Exec(query, tracker.Title, tracker.Notes, tracker.ID, tracker.UserID)
	if err != nil {
		fmt.Println("update error:", err)
		return err
	}

	return nil
}
