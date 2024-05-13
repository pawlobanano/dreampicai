package types

import "context"

// Config is a struct which encapsulates .env file variables.
type Config struct {
	DbHost         string
	DbName         string
	DbPassword     string
	DbPort         string
	DbUser         string
	Environment    string
	HttpListenAddr string
	SubabaseSecret string
	SupabaseUrl    string
}

// Logger interface for logging.
type Logger interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
}

// User.
type userKey string

const UserContextKey userKey = "user"

type AuthenticatedUser struct {
	Email      string
	IsLoggedIn bool
}
