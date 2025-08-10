package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	// how many connections can be opened
	DB.SetMaxOpenConns(10)
	// how many connections are open at all time even then not used
	DB.SetMaxIdleConns(5)

	fmt.Print("\n#-#-#  Connected to database  #-#-# \n\n")

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Could not create users table", err)
	}

	createNotesTable := `
	CREATE TABLE IF NOT EXISTS notes (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	user_id INTEGER NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createNotesTable)
	if err != nil {
		log.Fatal("Could not create notes table", err)
	}

	createTasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY,
	parent_id INTEGER,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	status TEXT NOT NULL,
	is_repeat BOOLEAN NOT NULL,
	interval INTEGER,
	notes TEXT,
	due_date TIMESTAMP,
	exec_at TIMESTAMP,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	user_id INTEGER NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createTasksTable)
	if err != nil {
		log.Fatal("Could not create tasks table", err)
	}

	createTrackerTable := `
	CREATE TABLE IF NOT EXISTS tracker (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	notes TEXT NOT NULL,
	start_date TIMESTAMP,
	best_streak INTEGER,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	user_id INTEGER NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createTrackerTable)
	if err != nil {
		log.Fatal("Could not create tracker table", err)
	}

}
