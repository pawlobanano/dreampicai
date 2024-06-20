package types

import (
	"context"
)

// Server is a struct which encapsulates the Config and Logger.
type Server struct {
	Config Config
	Logger Logger
}

// Config is a struct which encapsulates .env file variables.
type Config struct {
	DbHost                string
	DbName                string
	DbPassword            string
	DbPort                string
	DbSource              string
	DbUser                string
	Environment           string
	HttpListenAddr        string
	MigrationURL          string
	SessionAccessTokenKey string
	SessionSecret         string
	SessionUserKey        string
	SupabaseSecret        string
	SupabaseUrl           string
}

// contextKey is a custom type for context key.
type ContextKey string

// Logger interface for logging.
type Logger interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
}
