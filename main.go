package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abdillahzakkie/amuse-finance-backend/helpers"
	"github.com/abdillahzakkie/amuse-finance-backend/routes"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	// load .env file
	helpers.LoadEnv()
	
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	port := fmt.Sprintf("localhost:%s", os.Getenv("PORT"))

	setupRoutes(app)

	log.Println(fmt.Sprintln("Server listening on port:", port))
	log.Fatal(app.Listen(port))
}

func setupRoutes(app *fiber.App) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Amuse Finance Metrics Page"}))

	v1 := app.Group("/api/v1")
	v1.Post("/users/new", routes.CreateNewUser)
	v1.Get("/users/:id", routes.GetUserById)
}