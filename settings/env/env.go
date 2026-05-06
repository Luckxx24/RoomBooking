package env

import (
	"os"
	"strconv"
)

func GetString(Key, falback string) string {
	key, ok := os.LookupEnv(Key)

	if !ok {
		return falback
	}
	return key
}

func GetInt(Key string, fallback int) int {
	key, ok := os.LookupEnv(Key)

	if !ok {
		return fallback
	}

	keyint, err := strconv.Atoi(key)

	if err != nil {
		return fallback
	}

	return keyint
}
