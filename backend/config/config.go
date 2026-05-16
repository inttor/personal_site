package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	JWTExpireHours string
	ServerPort     string
	UploadDir      string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		DBHost:         getEnv("DB_HOST", "127.0.0.1"),
		DBPort:         getEnv("DB_PORT", "3306"),
		DBUser:         getEnv("DB_USER", "root"),
		DBPassword:     getEnv("DB_PASSWORD", "djldjl023491."),
		DBName:         getEnv("DB_NAME", "personal_site"),
		JWTSecret:      getEnv("JWT_SECRET", "personal-site-jwt-secret-key"),
		JWTExpireHours: getEnv("JWT_EXPIRE_HOURS", "72"),
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		UploadDir:      getEnv("UPLOAD_DIR", "./uploads"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
