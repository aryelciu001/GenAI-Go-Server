package config

import (
	"fmt"
	"os"
)

type Config struct {
	ProjectID     string
	Location      string
	ModelName     string
	FirestoreDbID string
	ServerPort    string
	BucketName    string
}

func LoadConfig() *Config {
	return &Config{
		ProjectID:     mustGetEnv("PROJECT_ID"),
		Location:      mustGetEnv("LOCATION"),
		ModelName:     mustGetEnv("MODEL_NAME"),
		FirestoreDbID: mustGetEnv("FIRESTORE_DB_ID"),
		ServerPort:    mustGetEnv("SERVER_PORT"),
		BucketName:    mustGetEnv("BUCKET_NAME"),
	}

}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("environment variable %v is not defined", key))
	}
	return val
}
