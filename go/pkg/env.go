package pkg

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv(key string) string {
	godotenv.Load(".env")

	return os.Getenv(key)
}