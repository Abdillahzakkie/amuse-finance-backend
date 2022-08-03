package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	startMessage := fmt.Sprintf("http://localhost:%s", "8080")
	log.Fatal(app.Listen(startMessage))
}