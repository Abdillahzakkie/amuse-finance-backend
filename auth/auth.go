package auth

import (
	"errors"
	"os"
	"time"

	"github.com/abdillahzakkie/amuse-finance-backend/validators"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidJWT = errors.New("invalid token")
)

type customClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

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


func GenerateJwtToken(username string) (string, error) {
	claims := customClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		return "", validators.ErrInternalServerError
	}
	return signedToken, nil
}
