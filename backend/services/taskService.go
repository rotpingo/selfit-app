package services

import (
	"database/sql"
	"fmt"
	"selfit/database"
	"selfit/dto"
	"selfit/models"
	"selfit/utils"
	"time"
)

func GetAllTasks(userId int64) ([]models.Task, error) {
	query := "SELECT * FROM tasks WHERE user_id = $1"
	rows, err := database.DB.Query(query, userId)
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

func GetAllProgressTasks(userId int64) ([]dto.TaskResponseDTO, error) {
	query := "SELECT id, title, content, is_repeat, interval, notes, due_date FROM tasks WHERE status = $1 AND user_id = $2"
	rows, err := database.DB.Query(query, models.StatusProgress, userId)
	if err != nil {
		fmt.Println("error fetching:", err)
		return nil, err
	}
	defer rows.Close()

	var tasksDto []dto.TaskResponseDTO
	for rows.Next() {
		var task dto.TaskResponseDTO
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Content,
			&task.IsRepeat,
			&task.Interval,
			&task.Notes,
			&task.DueDate,
		)
		if err != nil {
			fmt.Println("error scanning data:", err)
			return nil, err
		}
		tasksDto = append(tasksDto, task)
	}
	return tasksDto, nil
}

func getTaskById(id int64, userId int64) (*models.Task, error) {

	query := `
		SELECT id, title, content, status, is_repeat, interval, notes, due_date, exec_at, created_at, updated_at, user_id
		FROM tasks
		WHERE id = $1
		AND user_id = $2
	`

	row := database.DB.QueryRow(query, id, userId)

	var task models.Task
	err := row.Scan(
		&task.ID,
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
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task not found")
		}
		return nil, err
	}

	return &task, nil
}

func CreateTask(task *models.Task) error {

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
	query := `
		UPDATE tasks
		SET title = $1, content = $2, is_repeat = $3, interval = $4, due_date = $5,   updated_at = $6
		WHERE id = $7
		AND user_id = $8
	`

	_, err := database.DB.Exec(query, task.Title, task.Content, task.IsRepeat, task.Interval, task.DueDate, task.UpdatedAt, task.ID, task.UserID)
	if err != nil {
		fmt.Println("update error:", err)
		return err
	}

	return nil
}

func DeleteTaskById(id int64, userId int64) error {

	query := `DELETE FROM tasks WHERE id = $1 AND user_id = $2`

	_, err := database.DB.Exec(query, id, userId)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}

func CompleteTaskById(id int64, taskDto dto.EndTaskDTO, userId int64) error {

	status := models.StatusDone
	now := time.Now()

	query := `
		UPDATE tasks
		SET status = $1,
			notes = $2,
			exec_at = $3,
			updated_at = $4
		WHERE id = $5
		AND user_id = $6
	`
	_, err := database.DB.Exec(query, status, taskDto.Notes, now, now, id, userId)
	if err != nil {
		fmt.Println("Abort error: ", err)
		return err
	}

	task, _ := getTaskById(id, userId)
	if task.IsRepeat {
		var newTask models.Task
		newTask.ParentID = task.ID
		newTask.Title = task.Title
		newTask.Content = task.Content
		newTask.IsRepeat = task.IsRepeat
		newTask.Interval = task.Interval
		newTask.Notes = taskDto.Notes
		newTask.DueDate = utils.AddDays(now, int(task.Interval))

		err = CreateTask(&newTask)
		if err != nil {
			return fmt.Errorf("Error creating new task after completion: %w", err)
		}
	}

	return nil
}

func AbortTaskById(id int64, taskDto dto.EndTaskDTO, userId int64) error {

	status := models.StatusAborted
	now := time.Now()

	query := `
		UPDATE tasks
		SET status = $1,
			notes = $2,
			exec_at = $3,
			updated_at = $4
		WHERE id = $5
		AND user_id = $6
	`
	_, err := database.DB.Exec(query, status, taskDto.Notes, now, now, id, userId)
	if err != nil {
		fmt.Println("Abort error: ", err)
		return err
	}

	task, _ := getTaskById(id, userId)
	if task.IsRepeat {
		var newTask models.Task
		newTask.ParentID = task.ID
		newTask.Title = task.Title
		newTask.Content = task.Content
		newTask.IsRepeat = task.IsRepeat
		newTask.Interval = task.Interval
		newTask.Notes = taskDto.Notes
		newTask.DueDate = utils.AddDays(now, int(task.Interval))

		err = CreateTask(&newTask)
		if err != nil {
			return fmt.Errorf("Error creating new task after abort: %w", err)
		}
	}

	return nil
}
