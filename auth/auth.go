package auth

import (
	"errors"
	"os"
	"time"

	"github.com/abdillahzakkie/amuse_finance_backend/validators"
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

func Authenticate(jwtToken string) (username string, err error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SIGNING_KEY")), nil 
		},
	)
	if err != nil {
		return "", ErrInvalidJWT
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return "", ErrInvalidJWT
	}

	if claims.ExpiresAt.UTC().Unix() < time.Now().UTC().Unix() {
		return "", ErrInvalidJWT
	}
	username = claims.Username
	return username, nil
}