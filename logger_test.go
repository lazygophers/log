package log

import (
	"fmt"
	"testing"
)

func TestLoggerSetLevel(t *testing.T) {
	logger := newLogger()
	for _, level := range []Level{TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel} {
		t.Run(fmt.Sprintf("level:%d", level), func(t *testing.T) {
			logger.SetLevel(level)
			if logger.level != level {
				t.Errorf("SetLevel(%v) = %v, want %v", level, logger.level, level)
			}
		})
	}
}

func TestLoggerSetOutput(t *testing.T) {
	logger := newLogger()
	logger.SetOutput()
	if logger.out != nil {
		t.Errorf("SetOutput(nil) should clear out")
	}
}
