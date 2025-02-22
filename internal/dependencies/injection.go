package dependencies

import (
	"errors"
)

// injects all dependencies that should be
// present at runtime
//
// INFO: expects .env file to be loaded in
// already
func Inject() error {
	// 1. db dbClient
	dbClient, err := GetSupabaseClient()
	if err != nil {
		return err
	}
	if dbClient == nil {
		return errors.New("supabase client is nil")
	}
	// 2. openai client
	openaiClient := GetOpenAIClient()
	if openaiClient == nil {
		return errors.New("openai client is nil")
	}
	return nil
}
