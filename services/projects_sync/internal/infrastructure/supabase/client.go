package supabase

import (
	"fmt"
	"os"

	sb "github.com/supabase-community/supabase-go"
)

func NewSupabaseClient() (*sb.Client, error) {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	if url == "" || key == "" {
		return nil, fmt.Errorf("No define supabase credentials")
	}

	return sb.NewClient(url, key, nil)
}
