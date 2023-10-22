package environment

import (
	"os"
)

func EnvVariable(key string) string {
	os.Setenv("jwtKey", "12321312asdasdasdaq1231=-=;;/'/...;")

	return os.Getenv(key)
}
