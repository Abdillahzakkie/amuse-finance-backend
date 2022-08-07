package routes

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/abdillahzakkie/amuse-finance-backend/helpers"
	"github.com/abdillahzakkie/amuse-finance-backend/models"
	"github.com/abdillahzakkie/amuse-finance-backend/validators"
	"github.com/gofiber/fiber/v2"
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

func CreateNewUser(c *fiber.Ctx) error {
	var user models.User

	// Parsing the body of the request and assigning it to the reqBody variable.
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validators.ErrInvalidData,
		})
	}

	// Insert new user into the database.
	// check if there is an error in the database. If there is an error, 
	// it will return the error message.
	if err := us.CreateNewUser(&user); err != nil {
		switch {
			case err == validators.ErrInternalServerError:
				return helpers.RespondWithError(c, fiber.StatusInternalServerError, err)
			default:
				return helpers.RespondWithError(c, fiber.StatusBadRequest, err)
		}
	}

	return c.Status(fiber.StatusAccepted).JSON(user)
}

// It gets a user by ID from the database and returns it as JSON
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return helpers.RespondWithError(c, fiber.StatusNotFound, validators.ErrInvalidID)
	}
	user := models.User{
		ID: uint(userId),
	}
	
	// Get user by ID from the database
	// If user does not exist returns an ErrUserNotFound
	// if other error is thrown it'll return ErrInternalServerError
	if err := us.GetUserById(&user); err != nil {
		switch {
			case err == validators.ErrInternalServerError:
				return helpers.RespondWithError(c, fiber.StatusInternalServerError, err)
			case err == models.ErrUserNotFound:
				return helpers.RespondWithError(c, fiber.StatusNotFound, err)
			default:
				return helpers.RespondWithError(c, fiber.StatusBadRequest, err)
		}
	}
	
	return c.Status(fiber.StatusOK).JSON(user)
}