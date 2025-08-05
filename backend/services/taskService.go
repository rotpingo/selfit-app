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

func CreateTask(task *models.Task) error {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	// TODO: implement User
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

func UpdateTask(task *models.Task) error {
	task.UpdatedAt = time.Now()
	fmt.Println("task: ", task)
	query := `
		UPDATE tasks
		SET title = $1, content = $2, is_repeat = $3, interval = $4, due_date = $5,   updated_at = $6
		WHERE id = $7
	`

	_, err := database.DB.Exec(query, task.Title, task.Content, task.IsRepeat, task.Interval, task.DueDate, task.UpdatedAt, task.ID)
	if err != nil {
		fmt.Println("update error:", err)
		return err
	}

	return nil
}

func DeleteTaskById(id int) error {

	query := `DELETE FROM tasks WHERE id = $1`

	_, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
