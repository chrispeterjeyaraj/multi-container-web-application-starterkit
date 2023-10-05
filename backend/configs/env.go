package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	envFilePath := filepath.Join("/home/chris/github_repos/multi-container-web-application-starterkit/backend/", ".env")
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGOURI")
}
