package constants

import (
	"fmt"
	"os"
)

var PROJECT_ID = mustGetEnv("PROJECT_ID")
var LOCATION = mustGetEnv("LOCATION")
var MODEL_NAME = mustGetEnv("MODEL_NAME")
var FIRESTORE_DB_ID = mustGetEnv("FIRESTORE_DB_ID")
var SERVER_PORT = mustGetEnv("SERVER_PORT")
var BUCKET_NAME = mustGetEnv("BUCKET_NAME")

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("environment variable %v is not defined", key))
	}
	return val
}
