package supa

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var client *supabase.Client

func InitSupabase() error {
	apiURL := "https://fhhgthkemcgztpekjqyp.supabase.co"
	err := godotenv.Load()
	if err != nil {
		return err
	}

	apiKey := os.Getenv("SUPABASE_SERVICE_KEY")
	client, err := supabase.NewClient(apiURL, apiKey, nil)
	if err != nil {
		return err
	}

	_ = client
	return nil
}
