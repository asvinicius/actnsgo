package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DatabaseConfig
	JWT JWTConfig
}

type AppConfig struct {
	Name string
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret     string
	Expiration string
}

func Load() (Config, error) {

	err := godotenv.Load()

	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		App: AppConfig{
			Name: os.Getenv("APP_NAME"),
			Port: os.Getenv("APP_PORT"),
			Env:  os.Getenv("APP_ENV"),
		},
		DB: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		JWT: JWTConfig{
			Secret:     os.Getenv("JWT_SECRET"),
			Expiration: os.Getenv("JWT_EXPIRATION"),
		},
	}

	return cfg, nil
}
