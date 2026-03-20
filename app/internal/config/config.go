package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
    }
}



type FilePathConfig struct {
	FilePathOnDisk string
	FilePathOnDiskWithToken string
}

type Config struct {
	FilePath FilePathConfig
	FilePathToken FilePathConfig
}

func New() *Config {
	return &Config{
		FilePath: FilePathConfig{
			FilePathOnDisk: getEnv("FILE_PATH", ""),
		},
		FilePathToken: FilePathConfig{
			FilePathOnDiskWithToken: getEnv("FILE_PATH_TOKEN", ""),
		},
	}
}



func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}