package utils

import (
	"os"
	"strconv"
)

func GetStringEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	} else {
		return defaultValue
	}
}

func GetBoolEnv(key string, defaultValue bool) bool {
	value, ok := os.LookupEnv(key)
	if ok {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		}
		return boolValue
	} else {
		return defaultValue
	}
}
