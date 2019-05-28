package env

import "os"

// GetEnv tries to lookup the specified key from the OS environment variables. If not found, the fallback is returned.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
