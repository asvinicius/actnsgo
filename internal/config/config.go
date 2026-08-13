package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App     AppConfig
	DB      DatabaseConfig
	JWT     JWTConfig
	Cartola CartolaConfig
	Backup  BackupConfig
	Upload  UploadConfig
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

type CartolaConfig struct {
	CartolaURL string
}

type BackupConfig struct {
	Dir string
}

type UploadConfig struct {
	BankLogoDir   string
	AttachmentDir string
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
		Cartola: CartolaConfig{
			CartolaURL: os.Getenv("CARTOLA_BASE_URL"),
		},
		Backup: BackupConfig{
			Dir: os.Getenv("BACKUP_DIR"),
		},
		Upload: UploadConfig{
			BankLogoDir:   os.Getenv("UPLOAD_BANK_LOGO_DIR"),
			AttachmentDir: os.Getenv("ATTACHMENT_DIR"),
		},
	}

	return cfg, nil
}
