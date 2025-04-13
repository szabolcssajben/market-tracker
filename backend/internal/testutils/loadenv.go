package testutils

import (
	"testing"

	"github.com/joho/godotenv"
)

func LoadEnv(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal("Could not load .env file:", err)
	}
}
