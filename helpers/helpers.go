package helpers

import (
	"log"
	"os"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var rootPath string

func init() {
	rootDir := "amuse_finance_backend"
	re := regexp.MustCompile(`^(.*` + rootDir + `)`)
	cwd, _ := os.Getwd()
	rootPath = string(re.Find([]byte(cwd)))
}

// It takes the current working directory, finds the first occurrence of the string "amuse_finance_backend" in it, and
// then loads the .env file in that directory
func LoadEnv() {
	err := godotenv.Load(rootPath + `/.env`)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

// It takes a context, a status code, and a message, and returns an error
func RespondWithError(c *fiber.Ctx, status int, message error) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message.Error(),
	})
}