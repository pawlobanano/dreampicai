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
