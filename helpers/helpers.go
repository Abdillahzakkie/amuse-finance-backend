package helpers

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// It takes the current working directory, finds the first occurrence of the string "amuse-finance" in it, and
// then loads the .env file in that directory
func LoadEnv() {
	re := regexp.MustCompile(`^(.*` + "amuse-finance" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}