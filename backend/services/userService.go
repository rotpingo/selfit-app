package services

import (
	"errors"
	"fmt"
	"selfit/database"
	"selfit/models"
	"selfit/utils"
	"time"
)

func CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO users(email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	err = database.DB.QueryRow(
		query,
		user.Email,
		hashedPassword,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)

	if err != nil {
		fmt.Println("insert error:", err)
		return err
	}

	return nil
}

func ValidateUser(user *models.User) error {

	query := `
		SELECT password
		FROM users
		WHERE email = $1
	`

	row := database.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
