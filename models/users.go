package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/abdillahzakkie/amuse-finance-backend/database"
	"github.com/abdillahzakkie/amuse-finance-backend/validators"
)

var (
	ErrUserAlreadyExists	= errors.New("user already exists")
	ErrUserNotFound 		= errors.New("user not found")
)

// A User is a struct that has an ID, Username, Email, Biography, Password, CreatedAt, and UpdatedAt.
// @property {uint} ID - The user's ID.
// @property {string} Username - The username of the user.
// @property {string} Email - The email address of the user.
// @property {string} Biography - A short description of the user.
// @property {string} Password - This is the password of the user. We don't want to return this to the
// client, so we add the `-` tag to the property.
// @property CreatedAt - The time the user was created
// @property UpdatedAt - The time the user was last updated.
type User struct {
	ID 				uint 		`json:"id"`
	Username		string    	`json:"username"`
	Email			string    	`json:"email"`
	Biography		string 		`json:"biography"`
	Password  		string    	`json:"-"`
	CreatedAt 		time.Time 	`json:"-"`
	UpdatedAt 		time.Time 	`json:"-"`
}

// `UserService` is an interface that has a method called `CreateNewUser` that takes a pointer to a
// `User` and returns an error.
// @property {error} CreateNewUser - This is the method that will be used to create a new user.
type UserService interface {
	destructiveReset() error
}

// userDB is a struct type that implements the UserService interface.
// @property db - This is the database connection.
var _ UserService = &userDB{}
type userDB struct {
	db *sql.DB
}
// NewUserService create a new instance of UserService
// returns the newly created UserService instance or
// ErrInternalServerError if other error is encountered.
func NewUserService(psqlInfo string) (UserService, error) {
	db, err := database.NewClient(psqlInfo)
	if err != nil {
		return nil, validators.ErrInternalServerError
	}
	ud := userDB{
		db: db,
	}
	return &ud, nil
}

// Close closes the database connection.
// returns ErrInternalServerError if other error is encountered.
func (db *userDB) Close() error {
	if err := db.db.Close(); err != nil {
		return validators.ErrInternalServerError
	}
	return nil
}