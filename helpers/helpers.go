package helpers

import (
	"log"
	"os"
	"path/filepath"
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

// It reads the file at the given path and returns the content as a byte array
func GetFileContent(path string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(rootPath, path))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// It takes a context, a status code, and a message, and returns an error
func RespondWithError(c *fiber.Ctx, status int, message error) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message.Error(),
	})
}