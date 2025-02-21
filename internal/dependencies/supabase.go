package dependencies

import (
	"os"
	"sync"

	"github.com/supabase-community/supabase-go"
)

var (
	supabaseClient *supabase.Client
	once           sync.Once
)

// returns a singleton instance of
// supabase.Client
func GetSupabaseClient() (*supabase.Client, error) {
	url := os.Getenv("DB_URL")
	key := os.Getenv("DB_KEY")
	client, err := supabase.NewClient(url, key, &supabase.ClientOptions{})
	if err != nil {
		return nil, err
	}
	return client, nil
}
