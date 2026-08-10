// Package greeting reads a greeting from the environment.
package greeting

import "os"

func Greeting() string {
	if value := os.Getenv("GOSTLINGS_GREETING"); value != "" {
		return value
	}
	return "hello"
}
