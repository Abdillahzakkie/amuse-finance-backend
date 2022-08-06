package database

import (
	"fmt"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"amuse_finance_dev",
	)
	db, err := NewClient(psqlInfo)
	if err != nil {
		t.Error(err)
	}
	
	if db == nil {
		t.Error("Database connection failed")
	}
}