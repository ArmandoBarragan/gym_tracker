package conf

import (
	"fmt"
	"os"
	"strconv"
)

func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Environment variable %s not found", key))
	}
	return val
}

func getBoolEnv(key string) bool {
	var environmentValue string = getStrEnv(key)
	convertedValue, err := strconv.ParseBool(environmentValue)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s could not be converted to bool", key))
	}
	return convertedValue
}

type Settings struct {
	SECRET_KEY string
	DB_DSN     string
	PORT       string
	DEBUG      bool
	WHITELIST  []string
}

func InitSettings() *Settings {
	var dbDSN string = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		getStrEnv("POSTGRES_HOST"),
		getStrEnv("POSTGRES_USER"),
		getStrEnv("POSTGRES_PASSWORD"),
		getStrEnv("POSTGRES_NAME"),
		getStrEnv("POSTGRES_PORT"),
		getStrEnv("POSTGRES_SSL_MODE"),
		getStrEnv("POSTGRES_TIMEZONE"),
	)
	return &Settings{
		SECRET_KEY: getStrEnv("SECRET_KEY"),
		DB_DSN:     dbDSN,
		PORT:       getStrEnv("PORT"),
		DEBUG:      getBoolEnv("DEBUG"),
	}
}
