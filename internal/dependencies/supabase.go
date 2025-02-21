package dependencies

import (
	"os"
	"sync"

	"github.com/supabase-community/supabase-go"
)

var (
	supabaseClient *supabase.Client
	err            error
	once           sync.Once
)

// returns a singleton instance of
// supabase.Client
func GetSupabaseClient() (*supabase.Client, error) {
	once.Do(func() {
		url := os.Getenv("DB_URL")
		key := os.Getenv("DB_KEY")
		supabaseClient, err = supabase.NewClient(url, key, &supabase.ClientOptions{})
	})
	return supabaseClient, err
}
