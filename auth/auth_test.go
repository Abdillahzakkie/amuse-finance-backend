package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "test_example12345"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Error(err)
	}
	if hashedPassword == "" {
		t.Error("Hashed password is empty")
	}

	if password == hashedPassword {
		t.Error("Expected password to be different from hashed password")
	}

	// checks the returned hash password length is 60
	if len(hashedPassword) != 60 {
		t.Errorf("Expected password hash length to be 60, got %d", len(hashedPassword))
	}
}