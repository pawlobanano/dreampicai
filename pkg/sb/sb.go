package sb

import (
	"os"

	"github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func InitSupabaseClient() error {
	sbUrl := os.Getenv("SUPABASE_URL")
	sbSecret := os.Getenv("SUPABASE_SECRET")
	Client = supabase.CreateClient(sbUrl, sbSecret)

	return nil
}
