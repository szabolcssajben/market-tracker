package testutils

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func LoadEnv(t *testing.T) {
	if os.Getenv("CI") != "true" {
		err := godotenv.Load("../../.env")
		if err != nil {
			t.Fatal("Could not load .env file:", err)
		}
	}
}
