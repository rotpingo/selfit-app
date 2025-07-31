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

	createNotesTable := `
	CREATE TABLE IF NOT EXISTS notes (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	user_id INTEGER
	)
	`

	_, err := DB.Exec(createNotesTable)
	if err != nil {
		log.Fatal("Could not create notes table", err)
	}

}
