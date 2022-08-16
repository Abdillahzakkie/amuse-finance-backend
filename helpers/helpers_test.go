package helpers

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv()
	if os.Getenv("PORT") == "" {
		t.Error("Expected PORT to be not empty")
	}

	if os.Getenv("HOST_NAME") == "" {
		t.Error("Expected HOST_NAME to be not empty")
	}

	if os.Getenv("POSTGRES_USER") == "" {
		t.Error("Expected POSTGRES_USER to be not empty")
	}

	if os.Getenv("POSTGRES_PASSWORD") == "" {
		t.Error("Expected POSTGRES_PASSWORD to be not empty")
	}

	if os.Getenv("DB_NAME") == "" {
		t.Error("Expected DB_NAME to be not empty")
	}

	if os.Getenv("POSTGRES_PORT") == "" {
		t.Error("Expected POSTGRES_PORT to be not empty")
	}
}