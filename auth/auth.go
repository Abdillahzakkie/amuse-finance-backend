package auth

import (
	"os"

	"github.com/abdillahzakkie/amuse-finance-backend/validators"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashed the provided password
// returns the hashed password or ErrInvalidPassword if an error is encountered.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password + os.Getenv("PASSWORD_HASH_SECRET")), 
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", validators.ErrInvalidPassword
	}
	return string(hashedPassword), nil
}