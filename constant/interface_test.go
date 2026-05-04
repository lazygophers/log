package constant

import (
	"testing"
)

// Mock implementations for interface testing
type mockFormat struct{}

func (m mockFormat) Format(entry interface{}) []byte {
	return []byte("test")
}

type mockWriter struct{}

func (m mockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func TestHookFunc(t *testing.T) {
	t.Run("HookFunc_OnWrite_executes_function", func(t *testing.T) {
		called := false
		hook := HookFunc(func(entry interface{}) interface{} {
			called = true
			return entry
		})

		result := hook.OnWrite("test")

		if !called {
			t.Error("HookFunc should execute the function")
		}
		if result != "test" {
			t.Error("HookFunc should return the entry")
		}
	})

	t.Run("HookFunc_OnWrite_modifies_entry", func(t *testing.T) {
		hook := HookFunc(func(entry interface{}) interface{} {
			return "modified"
		})

		result := hook.OnWrite("test")

		if result != "modified" {
			t.Error("HookFunc should return modified entry")
		}
	})

	t.Run("HookFunc_OnWrite_filters_entry", func(t *testing.T) {
		hook := HookFunc(func(entry interface{}) interface{} {
			return nil // Filter out
		})

		result := hook.OnWrite("test")

		if result != nil {
			t.Error("HookFunc should return nil to filter")
		}
	})

	t.Run("HookFunc_OnWrite_with_nil_entry", func(t *testing.T) {
		hook := HookFunc(func(entry interface{}) interface{} {
			if entry == nil {
				return "nil handled"
			}
			return entry
		})

		result := hook.OnWrite(nil)

		if result != "nil handled" {
			t.Error("HookFunc should handle nil entry")
		}
	})
}

func TestHookInterface(t *testing.T) {
	t.Run("HookFunc_implements_Hook", func(t *testing.T) {
		var _ Hook = HookFunc(nil) // Compile-time check
	})
}

func TestFormatInterface(t *testing.T) {
	t.Run("Format_type_exists", func(t *testing.T) {
		var _ Format = mockFormat{}
	})
}

func TestWriterInterface(t *testing.T) {
	t.Run("Writer_type_exists", func(t *testing.T) {
		var _ Writer = mockWriter{}
	})
}
