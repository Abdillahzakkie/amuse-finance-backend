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

	if os.Getenv("DB_USER") == "" {
		t.Error("Expected DB_USER to be not empty")
	}

	if os.Getenv("DB_NAME") == "" {
		t.Error("Expected DB_NAME to be not empty")
	}

	if os.Getenv("DB_PORT") == "" {
		t.Error("Expected DB_PORT to be not empty")
	}
}