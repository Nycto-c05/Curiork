package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

// better error handling
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback //, fmt.Errorf("env %s not set", key)
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback //, fmt.Errorf("env %s invalid int: %v", key, err)
	}

	return valAsInt
}
