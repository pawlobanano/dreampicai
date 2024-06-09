package sb

import (
	"dreampicai/types"

	"github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func InitSupabaseClient(s types.Server) error {
	Client = supabase.CreateClient(
		s.Config.SupabaseUrl,
		s.Config.SupabaseSecret,
	)

	return nil
}
