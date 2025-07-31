package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
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

	fmt.Print("\n#-#-#  Connected to database  #-#-# \n\n")
}
