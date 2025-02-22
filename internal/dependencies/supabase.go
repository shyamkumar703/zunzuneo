package dependencies

import (
	"os"
	"sync"

	"github.com/supabase-community/supabase-go"
)

var (
	supabaseClient  *supabase.Client
	supabaseInitErr error
	supabaseOnce    sync.Once
)

// returns a singleton instance of
// supabase.Client
func GetSupabaseClient() (*supabase.Client, error) {
	supabaseOnce.Do(func() {
		url := os.Getenv("DB_URL")
		key := os.Getenv("DB_KEY")
		supabaseClient, supabaseInitErr = supabase.NewClient(url, key, &supabase.ClientOptions{})
	})
	return supabaseClient, supabaseInitErr
}
