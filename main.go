package main

import (
	"zunzuneo/internal/dependencies"

	"github.com/joho/godotenv"
)

func main() {
	envError := godotenv.Load(".env")
	if envError != nil {
		panic(envError)
	}
	depErr := dependencies.Inject()
	if depErr != nil {
		panic(depErr)
	}
}
