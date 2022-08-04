package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// NewClient takes the environment variables and uses them to connect to the database
func NewClient() (*sql.DB, error) {
	host := os.Getenv("HOST_NAME")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	
	fmt.Println("Connected to database!")
	return db, nil
}