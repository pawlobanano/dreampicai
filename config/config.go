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
	if err := checkEnvVar("DB_SOURCE"); err != nil {
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
	if err := checkEnvVar("MIGRATION_URL"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("SESSION_ACCESS_TOKEN_KEY"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("SESSION_SECRET"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("SESSION_USER_KEY"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("SUPABASE_SECRET"); err != nil {
		return types.Config{}, err
	}
	if err := checkEnvVar("SUPABASE_URL"); err != nil {
		return types.Config{}, err
	}

	config := types.Config{
		DbHost:                os.Getenv("DB_HOST"),
		DbName:                os.Getenv("DB_NAME"),
		DbPassword:            os.Getenv("DB_PASSWORD"),
		DbPort:                os.Getenv("DB_PORT"),
		DbSource:              os.Getenv("DB_SOURCE"),
		DbUser:                os.Getenv("DB_USER"),
		Environment:           os.Getenv("ENVIRONMENT"),
		HttpListenAddr:        os.Getenv("HTTP_LISTEN_ADDR"),
		MigrationURL:          os.Getenv("MIGRATION_URL"),
		SessionAccessTokenKey: os.Getenv("SESSION_ACCESS_TOKEN_KEY"),
		SessionSecret:         os.Getenv("SESSION_SECRET"),
		SessionUserKey:        os.Getenv("SESSION_USER_KEY"),
		SupabaseSecret:        os.Getenv("SUPABASE_SECRET"),
		SupabaseUrl:           os.Getenv("SUPABASE_URL"),
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
