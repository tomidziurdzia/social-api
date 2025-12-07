package env

import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func GetInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return i
}