package supa

import (
	"os"

	"github.com/supabase-community/supabase-go"
)

var client *supabase.Client

func InitSupabase() error {
	apiURL := "https://fhhgthkemcgztpekjqyp.supabase.co"
	apiKey, err := os.ReadFile("/run/secrets/supabase-service-key")
	if err != nil {
		return err
	}

	client, err := supabase.NewClient(apiURL, string(apiKey), nil)
	if err != nil {
		return err
	}

	_ = client
	return nil
}
