package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abdillahzakkie/amuse-finance-backend/helpers"
	"github.com/abdillahzakkie/amuse-finance-backend/routes"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	config(app, port)
	setupRoutes(app)

	log.Println(fmt.Sprintln("Server listening on port:", port))
	log.Fatal(app.Listen(port))
}

func config(app *fiber.App, port string) {
	app.Use(logger.New())
	
	app.Use(cors.New(cors.Config{
		// AllowOrigins: "*",
		AllowOrigins: port,
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	// app.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token",
	// 	CookieName:     "csrf_",
	// 	CookieSameSite: "Strict",
	// 	CookieHTTPOnly: true,
	// 	Expiration:     time.Hour * 24,
	// }))

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: os.Getenv("COOKIES_PASSPHRASE"),
	}))
}

func setupRoutes(app *fiber.App) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Amuse Finance Metrics Page"}))

	v1 := app.Group("/api/v1")
	v1.Post("/users/new", routes.CreateNewUser)
	v1.Get("/users/:id", routes.GetUserById)
}