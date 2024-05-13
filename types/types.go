package types

// Config is a struct which encapsulates .env file variables.
type Config struct {
	DbHost         string
	DbName         string
	DbPassword     string
	DbPort         string
	DbUser         string
	HttpListenAddr string
	SubabaseSecret string
	SupabaseUrl    string
}

// Logger interface for logging.
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// User.
type userKey string

const UserContextKey userKey = "user"

type AuthenticatedUser struct {
	Email      string
	IsLoggedIn bool
}
