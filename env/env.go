package env

import "os"

func Find(key string, defaultValue string) string {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}

	return val
}
