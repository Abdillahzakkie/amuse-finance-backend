package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abdillahzakkie/amuse-finance-backend/helpers"
	"github.com/abdillahzakkie/amuse-finance-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// load .env file
	helpers.LoadEnv()
	app := fiber.New()

	app.Use(logger.New())
	setupRoutes(app)

	startMessage := fmt.Sprintf("localhost:%s", os.Getenv("PORT"))
	log.Println(startMessage)
	log.Fatal(app.Listen(startMessage))
}

func setupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Post("/users/new", routes.CreateNewUser)
	v1.Get("/users/:id", routes.GetUserById)
}