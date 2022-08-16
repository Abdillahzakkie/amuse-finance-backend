package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/abdillahzakkie/amuse_finance_backend/auth"
	"github.com/abdillahzakkie/amuse_finance_backend/database"
	"github.com/abdillahzakkie/amuse_finance_backend/helpers"
	"github.com/abdillahzakkie/amuse_finance_backend/validators"
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
	Token			string 		`json:"-"`
	CreatedAt 		time.Time 	`json:"-"`
	UpdatedAt 		time.Time 	`json:"-"`
}

type UserForm struct {
	ID 				uint 		`json:"id"`
	Username		string    	`json:"username"`
	Email			string    	`json:"email"`
	Biography		string 		`json:"biography"`
	Password  		string    	`json:"password"`
}

// `UserService` is an interface that has a method called `CreateNewUser` that takes a pointer to a
// `User` and returns an error.
// @property {error} CreateNewUser - This is the method that will be used to create a new user.
type UserService interface {
	CreateNewUser(body *UserForm) (User,error)
	IsExistingUser(username string) bool
	GetUserById(id uint) (User, error)
	DeleteUserByUserId(id uint) (User, error)
	Login(username, password string) (User, error)

	Migrate() error
	destructiveReset() error
	Close() error
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

	// Migrate the database
	ud.Migrate()

	// clear all users in users table
	ud.destructiveReset()
	return &ud, nil
}

// A function that is used to migrate the database.
func (db *userDB) Migrate() error {
	path := "./models/SQL/user.sql"
	data, err := helpers.GetFileContent(path)
	if err != nil {
		log.Fatal(err)
	}
	query := string(data)
	_, err = db.db.Exec(query)
	if err != nil {
		return validators.ErrInternalServerError
	}
	return nil
}

// Close closes the database connection.
// returns ErrInternalServerError if other error is encountered.
func (db *userDB) Close() error {
	if err := db.db.Close(); err != nil {
		return validators.ErrInternalServerError
	}
	return nil
}

// destructiveReset clears all records in users table
// returns ErrInternalServerError if other error is encountered.
func (db *userDB) destructiveReset() error {
	query := "TRUNCATE TABLE users CASCADE"
	_, err := db.db.Exec(query)
	if err != nil {
		return validators.ErrInternalServerError
	}
	return nil
}

// CreateNewUser creates a new user in the database.
// returns ErrUserAlreadyExists if user is already existed,
// or ErrInternalServerError if other error is encountered.
func (db *userDB) CreateNewUser(body *UserForm) (User,error) {
	user := User{
		Username: body.Username,
		Email: body.Email,
		Biography: body.Biography,
		Password: body.Password,
	}

	// Checking if the user is valid.
	if err := validators.ValidateUser(body.Username, body.Email, body.Password); err != nil {
		return User{}, err
	}
	// checking if the user is already existing.
	if db.IsExistingUser(body.Username) {
		return User{}, ErrUserAlreadyExists
	}

	// hash password
	var err error
	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return User{},validators.ErrInternalServerError
	}

	// save user to DB
	query := `
		INSERT INTO users (username, email, biography, password) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	row := db.db.QueryRow(query, user.Username, user.Email, user.Biography, user.Password)
	if err := row.Scan(&user.ID); err != nil {
		return User{}, validators.ErrInternalServerError
	}

	// generate JWT token
	user.Token, err = auth.GenerateJwtToken(user.Username)
	if err != nil {
		return User{}, err
	}
	// clear password from memory
	body.Password = ""
	user.Password = ""
	return user, nil
}

// GetUserById gets user by ID,
// returns ErrUserNotFound if user is not found
// or ErrInternalServerError if other error is encountered
func (db *userDB) GetUserById(id uint) (User, error) {
	var user User
	query := `
		SELECT id, username, email FROM users
		WHERE (id = $1) AND deleted_at IS NULL
	`
	row := db.db.QueryRow(query, id)
	if err := row.Scan(&user.ID, &user.Username, &user.Email); err != nil {
		switch {
			case err == sql.ErrNoRows:
				return User{}, ErrUserNotFound
			default:
				return User{}, validators.ErrInternalServerError
		}
	}
	return user, nil
}

// IsExistingUser checks if the user is already existing in the database.
// returns true if user already existed, false otherwise.
func (db *userDB) IsExistingUser(username string) bool {
	_, err := db.getUserByUsername(username)
	if err != nil {
		switch {
			case err == ErrUserNotFound:
				return false
			}
	}
	return true
}

// GetUser gets user from the database
// returns ErrUserNotFound if user is not found
// or ErrInternalServerError if other error is encountered
func (db *userDB) getUserByUsername(username string) (User,error) {
	var user User
	query := `
		SELECT id, username, email, password, biography FROM users
		WHERE username =  $1 AND deleted_at IS NULL
	`
	row := db.db.QueryRow(query, username)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Biography); err != nil {
		switch {
			case err == sql.ErrNoRows:
				return User{}, ErrUserNotFound
			default:
				return User{}, validators.ErrInternalServerError
		}
	}
	return user, nil
}
