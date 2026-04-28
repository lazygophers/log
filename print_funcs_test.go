package log

import (
	"testing"
)

func TestDebug(t *testing.T) {
	SetLevel(DebugLevel)
	Debug("test debug message")
	Debug("multiple", "args", 123)
}

func TestDebugf(t *testing.T) {
	SetLevel(DebugLevel)
	Debugf("formatted: %s", "message")
	Debugf("number: %d", 42)
}
