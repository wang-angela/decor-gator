package utils

import (
	"os"
)

func GetEnvVar(key string) string {
	return os.Getenv(key)
}
