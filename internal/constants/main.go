package constants

import (
	"fmt"
	"os"
)

var PROJECT_ID = getEnvOrPanic("PROJECT_ID")
var LOCATION = getEnvOrPanic("LOCATION")
var MODEL_NAME = getEnvOrPanic("MODEL_NAME")
var FIRESTORE_DB_ID = getEnvOrPanic("FIRESTORE_DB_ID")
var SERVER_PORT = getEnvOrPanic("SERVER_PORT")

func getEnvOrPanic(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("environment variable %v is not defined", key))
	}
	return val
}
