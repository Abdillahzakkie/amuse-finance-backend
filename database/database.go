package database

import (
	"database/sql"
	"fmt"

	"github.com/abdillahzakkie/amuse_finance_backend/helpers"
	_ "github.com/lib/pq"
)

// NewClient takes the environment variables and uses them to connect to the database
func NewClient(psqlInfo string) (*sql.DB, error) {
	// load .env file to environment variables
	helpers.LoadEnv()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	
	fmt.Println("Connected to database!")
	return db, nil
}