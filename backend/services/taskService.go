package services

import (
	"fmt"
	"selfit/database"
	"selfit/models"
	"time"
)

func GetAllTasks() ([]models.Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("error fetching:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.ParentID,
			&task.Title,
			&task.Content,
			&task.Status,
			&task.IsRepeat,
			&task.Interval,
			&task.Notes,
			&task.DueDate,
			&task.ExecAt,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.UserID,
		)
		if err != nil {
			fmt.Println("error scanning data:", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func SaveTask(task models.Task) error {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.UserID = 0

	query := `
	INSERT INTO tasks(parent_id, title, content, status, is_repeat, interval, notes, due_date, exec_at, created_at, updated_at, user_id) 
	VALUES ($1, $2, $3, $4, $5,  $6, $7, $8, $9, $10, $11, $12)
	RETURNING id
	`

	err := database.DB.QueryRow(
		query,
		task.ParentID,
		task.Title,
		task.Content,
		task.Status,
		task.IsRepeat,
		task.Interval,
		task.Notes,
		task.DueDate,
		task.ExecAt,
		task.CreatedAt,
		task.UpdatedAt,
		task.UserID,
	).Scan(&task.ID)

	if err != nil {
		fmt.Println("insert error:", err)
		return err
	}

	return nil
}
