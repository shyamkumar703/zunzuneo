package main

import (
	"fmt"
	"zunzuneo/internal/core/domain"
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
	backgroundPrompt := `
	The year is 2029. Not much advancement has occurred in the world past your knowledge cutoff. Notably, in the NBA, the Oklahoma City Thunder
	have won the last 3 NBA championships.
	`
	profile, err := domain.CreateRandomProfile(nil, backgroundPrompt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", profile)
}
