package utils

import (
	"os"
)

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(key)
	}
	return value
}
