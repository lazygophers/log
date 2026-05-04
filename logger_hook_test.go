package log

import (
	"strings"
	"testing"

	"github.com/lazygophers/log/constant"
)

func TestHook_Basic(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	called := false
	hook := constant.HookFunc(func(entry interface{}) interface{} {
		called = true
		return entry
	})

	logger.AddHook(hook).Info("test")

	if !called {
		t.Error("hook was not called")
	}
}

func TestHook_Filter(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// Hook that returns nil to filter
	hook := constant.HookFunc(func(entry interface{}) interface{} {
		if e, ok := entry.(*Entry); ok && strings.Contains(e.Message, "skip") {
			return nil
		}
		return entry
	})

	loggerWithHook := logger.AddHook(hook)

	// Should log
	loggerWithHook.Info("normal message")

	// Should be filtered
	loggerWithHook.Info("skip this")
}

func TestHook_ModifyMessage(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// Hook that modifies message
	hook := constant.HookFunc(func(entry interface{}) interface{} {
		if e, ok := entry.(*Entry); ok {
			e.Message = "[MODIFIED] " + e.Message
		}
		return entry
	})

	logger.AddHook(hook).Info("test")

	// Message should be modified
}

func TestHook_Chain(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	count := 0
	hook1 := constant.HookFunc(func(entry interface{}) interface{} {
		count++
		return entry
	})

	hook2 := constant.HookFunc(func(entry interface{}) interface{} {
		count++
		return entry
	})

	logger.AddHooks(hook1, hook2).Info("test")

	if count != 2 {
		t.Errorf("expected 2 hooks called, got %d", count)
	}
}

func TestHook_RemoveHooks(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	called := false
	hook := constant.HookFunc(func(entry interface{}) interface{} {
		called = true
		return entry
	})

	logger.AddHook(hook)
	logger.RemoveHooks().Info("test")

	if called {
		t.Error("hook should not be called after removal")
	}
}
