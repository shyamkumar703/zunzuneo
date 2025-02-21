package main

import (
	"fmt"
	"zunzuneo/internal/dependencies"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	client, err := dependencies.GetSupabaseClient()
	if err != nil {
		panic(err)
	} else if client == nil {
		panic("client is nil")
	} else {
		fmt.Printf("everything is good!\n")
	}
}
