package services

import (
	"database/sql"
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
		SELECT id, password
		FROM users
		WHERE email = $1
	`

	row := database.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}

func GetUser(userId int64) (models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"

	var user models.User
	err := database.DB.QueryRow(query, userId).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("user not found")
		} else {
			fmt.Println("error fetching user data:", err)
		}
		return models.User{}, err
	}

	return user, nil
}
