package main

import (
	"zunzuneo/internal/dependencies"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	openAIClient := dependencies.GetOpenAIClient()
	if openAIClient == nil {
		panic("openai client is nil")
	}
	supabaseClient, err := dependencies.GetSupabaseClient()
	if err != nil {
		panic(err)
	} else if supabaseClient == nil {
		panic("client is nil")
	}
}
