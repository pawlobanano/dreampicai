package config

import (
	"dreampicai/types"
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvVars loads the configuration from the .env file and returns a Config and error if any.
func LoadEnvVars() (types.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return types.Config{}, err
	}

	if err := checkEnvVar("DB_HOST"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("DB_NAME"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("DB_PASSWORD"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("DB_PORT"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("DB_USER"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("ENVIRONMENT"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("HTTP_LISTEN_ADDR"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("SESSION_SECRET"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("SUPABASE_SECRET"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("SUPABASE_URL"); err != nil {
		return types.Config{}, err
	}

	config := types.Config{
		DbHost:         os.Getenv("DB_HOST"),
		DbName:         os.Getenv("DB_NAME"),
		DbPassword:     os.Getenv("DB_PASSWORD"),
		DbPort:         os.Getenv("DB_PORT"),
		DbUser:         os.Getenv("DB_USER"),
		Environment:    os.Getenv("ENVIRONMENT"),
		HttpListenAddr: os.Getenv("HTTP_LISTEN_ADDR"),
		SessionSecret:  os.Getenv("SESSION_SECRET"),
		SupabaseSecret: os.Getenv("SUPABASE_SECRET"),
		SupabaseUrl:    os.Getenv("SUPABASE_URL"),
	}

	return config, nil
}

// checkEnvVar checks if the environment variable is set and not empty.
func checkEnvVar(varName string) error {
	if varValue, ok := os.LookupEnv(varName); !ok {
		return errors.New(varName + " variable not found")
	} else if varValue == "" {
		return errors.New(varName + " variable found but empty")
	}
	return nil
}
