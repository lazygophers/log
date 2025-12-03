package log

import (
	"os"
	"strings"
)

// 通过环境变量设置日志级别
func init() {
	switch strings.ToLower(os.Getenv("APP_ENV")) {
	case "dev", "development":
		SetLevel(DebugLevel)
	case "test", "canary":
		SetLevel(DebugLevel)
	case "prod", "release", "production":
		SetLevel(InfoLevel)
	}
}
