package util

import "os"

// GetEnv - get environment variable value
func GetEnv(params ...string) string {
	value := os.Getenv(params[0])
	if value == "" && len(params) == 2 {
		return params[1]
	}
	return value
}
