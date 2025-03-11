package main

import (
	"fmt"
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
	response, err := dependencies.RequstLLM("this is a test", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is %s\n", *response)
}
