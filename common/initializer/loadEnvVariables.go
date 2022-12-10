package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(finishedInit chan bool) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	finishedInit <- true
}
