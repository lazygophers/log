package zap

import (
	"fmt"
	"io"
	"os"
	"sync"

	"go.uber.org/zap/zapcore"
)

type Hook struct {
	mu     sync.Mutex
	output io.Writer
}

func NewHook(w io.Writer) *Hook {
	if w == nil { w = os.Stdout }
	return &Hook{output: w}
}

func (h *Hook) Write(p []byte) (n int, err error) {
	h.mu.Lock(); defer h.mu.Unlock()
	return h.output.Write(p)
}

func (h *Hook) Sync() error {
	h.mu.Lock(); defer h.mu.Unlock()
	if s, ok := h.output.(interface{ Sync() error }); ok {
		return s.Sync()
	}
	return nil
}

type WriteSyncer struct { hook *Hook }

func NewWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	return &WriteSyncer{hook: NewHook(w)}
}

func (w *WriteSyncer) Write(p []byte) (n int, err error) {
	return w.hook.Write(p)
}

func (w *WriteSyncer) Sync() error {
	return w.hook.Sync()
}

type Logger struct {
	mu     sync.Mutex
	output io.Writer
}

func NewLogger(w io.Writer) *Logger {
	if w == nil { w = os.Stdout }
	return &Logger{output: w}
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.mu.Lock(); defer l.mu.Unlock()
	return l.output.Write(p)
}

func (l *Logger) Sync() error {
	l.mu.Lock(); defer l.mu.Unlock()
	if s, ok := l.output.(interface{ Sync() error }); ok {
		return s.Sync()
	}
	return nil
}

func (l *Logger) Log(level, msg string) {
	l.mu.Lock(); defer l.mu.Unlock()
	fmt.Fprintln(l.output, fmt.Sprintf("[%s] %s", level, msg))
}
