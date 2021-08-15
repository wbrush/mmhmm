package configuration

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/wbrush/go-common/config"
)

type (
	Config struct {
		Version string `json:"version"`
		BuiltAt string `json:"builtAt"`

		Host     string `json:"host"`
		Port     string `json:"port"`
		Environ  string `json:"environ"`
		LogLevel string `json:"logLevel`

		config.ServiceParams
		config.DbParams

		DbMigrationPath string `json:"db_migration_path"`
	}
)

var cfg Config

func InitConfig(commit, builtAt string) *Config {
	//  load space .env variables first if available
	filename := "./files/.env"
	if _, err := os.Stat(filename); err == nil {
		_ = godotenv.Load(filename)
	}

	cfg := &Config{
		ServiceParams: config.ServiceParams{
			Version: commit,
			BuiltAt: builtAt,

			Environment:  os.Getenv("ENVIRONMENT"),
			GlobalRegion: "",

			Host:     os.Getenv("HOST"),
			BaseUri:  "https://localhost",
			Port:     os.Getenv("PORT"),
			LogLevel: os.Getenv("LOG_LEVEL"),
		},

		DbParams: config.DbParams{
			Host:     os.Getenv("LOG_LEVEL"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DATABASE"),
			NumConns: "5",
		},

		DbMigrationPath: "./dao/postgres",
	}

	SetCurrentCfg(*cfg)

	return cfg
}

func GetCurrentCfg() Config {
	return cfg
}

func SetCurrentCfg(c Config) {
	cfg = c
}
