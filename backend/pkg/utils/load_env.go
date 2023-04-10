package utils

import (
	"os"
)

func GetEnvVar(key string) []byte {
	return []byte(os.Getenv(key))
}
