package log

import (
	"bytes"
	"strings"
	"testing"

	"github.com/lazygophers/log/constant"
)

func TestHooksFiltering(t *testing.T) {
	t.Run("Hook_filters_entry", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)

		// Hook that filters all messages
		filterHook := constant.HookFunc(func(entry interface{}) interface{} {
			return nil // Filter out
		})

		logger.AddHook(filterHook)
		logger.Info("filtered message")

		if buf.Len() != 0 {
			t.Error("filtered message should not be logged")
		}
	})

	t.Run("Hook_chain_with_filter", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)

		// First hook passes, second filters
		hook1 := constant.HookFunc(func(entry interface{}) interface{} {
			return entry
		})
		hook2 := constant.HookFunc(func(entry interface{}) interface{} {
			return nil
		})

		logger.AddHooks(hook1, hook2)
		logger.Info("test")

		if buf.Len() != 0 {
			t.Error("should be filtered by second hook")
		}
	})

	t.Run("Hook_modifies_entry", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)

		modifyHook := constant.HookFunc(func(entry interface{}) interface{} {
			if e, ok := entry.(*Entry); ok {
				e.Message = "[MODIFIED] " + e.Message
			}
			return entry
		})

		logger.AddHook(modifyHook)
		logger.Info("test")

		if !strings.Contains(buf.String(), "[MODIFIED]") {
			t.Error("message should be modified by hook")
		}
	})

	t.Run("Hook_with_fields_enrichment", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)

		enrichHook := constant.HookFunc(func(entry interface{}) interface{} {
			if e, ok := entry.(*Entry); ok {
				// Add field
				e.Fields = append(e.Fields, KV{Key: "enriched", Value: true})
			}
			return entry
		})

		logger.AddHook(enrichHook)
		logger.Info("test")

		if !strings.Contains(buf.String(), "enriched") {
			t.Error("should include enriched field")
		}
	})

	t.Run("ApplyHooks_with_nil_entry", func(t *testing.T) {
		logger := New()

		// Test applyHooks with nil entry (hook chain continues but returns nil)
		filterHook := constant.HookFunc(func(entry interface{}) interface{} {
			if entry == nil {
				return nil
			}
			return entry
		})

		logger.AddHook(filterHook)
		// This will call applyHooks with entry != nil
		logger.Info("test")
	})

	t.Run("RemoveHooks_clears_all", func(t *testing.T) {
		logger := New()

		called := false
		hook := constant.HookFunc(func(entry interface{}) interface{} {
			called = true
			return entry
		})

		logger.AddHook(hook)
		logger.RemoveHooks()

		var buf bytes.Buffer
		logger.SetOutput(&buf)
		logger.Info("test")

		if called {
			t.Error("hook should not be called after removal")
		}
	})
}

func TestLoggerMultipleHooks(t *testing.T) {
	t.Run("Multiple_hooks_execution_order", func(t *testing.T) {
		order := []string{}
		logger := New()

		hook1 := constant.HookFunc(func(entry interface{}) interface{} {
			order = append(order, "hook1")
			return entry
		})

		hook2 := constant.HookFunc(func(entry interface{}) interface{} {
			order = append(order, "hook2")
			return entry
		})

		hook3 := constant.HookFunc(func(entry interface{}) interface{} {
			order = append(order, "hook3")
			return entry
		})

		logger.AddHooks(hook1, hook2, hook3)
		logger.Info("test")

		if len(order) != 3 {
			t.Errorf("expected 3 hooks, got %d", len(order))
		}

		if order[0] != "hook1" || order[1] != "hook2" || order[2] != "hook3" {
			t.Error("hooks should execute in order")
		}
	})
}
