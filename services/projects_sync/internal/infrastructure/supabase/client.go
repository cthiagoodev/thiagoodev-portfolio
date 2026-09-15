package supabase

import sb "github.com/supabase-community/supabase-go"

func NewSupabaseClient() (*sb.Client, error) {
	url := ""
	key := ""

	return sb.NewClient(url, key, nil)
}
