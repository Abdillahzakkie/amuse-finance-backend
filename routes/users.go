package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/abdillahzakkie/amuse-finance-backend/helpers"
	"github.com/abdillahzakkie/amuse-finance-backend/models"
)

var us models.UserService

func init() {
	helpers.LoadEnv()
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		os.Getenv("HOST_NAME"), 
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"),
	)
	var err error
	us, err = models.NewUserService(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
}

