package log

import (
	"os"
	"strings"
)

// init sets the log level based on the APP_ENV environment variable.
func init() {
	switch strings.ToLower(os.Getenv("APP_ENV")) {
	case "dev", "development":
		SetLevel(TraceLevel)
	case "test", "canary":
		SetLevel(DebugLevel)
	case "prod", "release", "production":
		SetLevel(InfoLevel)
	}
}
