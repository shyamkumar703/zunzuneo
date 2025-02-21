package dependencies_test

import (
	"testing"
	"zunzuneo/internal/dependencies"

	"github.com/joho/godotenv"
)

func TestNewSupabaseClient(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("could not initialize environment variables, err=%s", err)
	}
	client, err := dependencies.GetSupabaseClient()
	if err != nil {
		t.Fatalf("could not initialize supabase client, err=%s", err)
	}
	if client == nil {
		t.Fatalf("initialization did not return an error, but client is nil")
	}
}

