package utils

import (
	"os"
)

const (
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
	HeaderUserAgent     = "User-Agent"
	HeaderAccept        = "Accept"
)

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(key)
	}
	return value
}
