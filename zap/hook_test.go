package zap

import (
	"bytes"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestNewHook(t *testing.T) {
	hook := NewHook(nil)
	if hook == nil { t.Fatal("nil") }
	if hook.output == nil { t.Error("output nil") }
}

func TestHook_Write(t *testing.T) {
	var buf bytes.Buffer
	hook := NewHook(&buf)
	n, err := hook.Write([]byte("test"))
	if err != nil { t.Fatal(err) }
	if n != 4 { t.Errorf("got %d", n) }
}

func TestWriteSyncer_Interface(t *testing.T) {
	var _ zapcore.WriteSyncer = NewWriteSyncer(nil)
}
