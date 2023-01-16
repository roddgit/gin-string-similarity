package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return EnvConfig()["MONGOURI"]
}

func EnvConfig() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	result := make(map[string]string)

	result["PORT"] = getEnv("PORT", "8080")
	result["DB_METHOD"] = getEnv("DB_METHOD", "MONGO")
	result["MONGOURI"] = getEnv("MONGOURI", "mongodb://root:password@localhost:27017")
	result["OR_HOST_DB"] = getEnv("OR_HOST_DB", "localhost")
	result["OR_USER_DB"] = getEnv("OR_USER_DB", "root")
	result["OR_PASSWORD_DB"] = getEnv("OR_PASSWORD_DB", "password")
	result["OR_SID"] = getEnv("OR_SID", "XE")
	result["OR_PORT_DB"] = getEnv("OR_PORT_DB", "1521")
	result["LOGS_FILES"] = getEnv("LOGS_FILES", "ENABLE")
	result["LOGS_FOLDER"] = getEnv("LOGS_FOLDER", "/logs")
	result["LOGS_ROTATION_MB"] = getEnv("LOGS_ROTATION_MB", "200")
	result["LOGS_RETENTION_DAYS"] = getEnv("LOGS_RETENTION_DAYS", "6")

	return result
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
