package config

import (
	"dreampicai/types"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvVars loads the configuration from the .env file and returns a Config and error.
func LoadEnvVars(log types.Logger) (*types.Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error("Loading .env file", err)
		return nil, err
	}

	checkEnvVar(log, "DB_HOST")
	checkEnvVar(log, "DB_NAME")
	checkEnvVar(log, "DB_PASSWORD")
	checkEnvVar(log, "DB_PORT")
	checkEnvVar(log, "DB_USER")
	checkEnvVar(log, "HTTP_LISTEN_ADDR")
	checkEnvVar(log, "SUPABASE_SECRET")
	checkEnvVar(log, "SUPABASE_URL")

	config := &types.Config{
		DbHost:         os.Getenv("DB_HOST"),
		DbName:         os.Getenv("DB_NAME"),
		DbPassword:     os.Getenv("DB_PASSWORD"),
		DbPort:         os.Getenv("DB_PORT"),
		DbUser:         os.Getenv("DB_USER"),
		HttpListenAddr: os.Getenv("HTTP_LISTEN_ADDR"),
		SubabaseSecret: os.Getenv("SUPABASE_SECRET"),
		SupabaseUrl:    os.Getenv("SUPABASE_URL"),
	}

	return config, nil
}

// checkEnvVar checks if the environment variable is set and not empty.
func checkEnvVar(log types.Logger, varName string) {
	if varValue, ok := os.LookupEnv(varName); !ok {
		log.Error(varName + " variable not found")
	} else if varValue == "" {
		log.Error(varName + " variable found but empty")
	}
}
