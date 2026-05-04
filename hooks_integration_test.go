package log

import (
	"strings"
	"testing"

	"github.com/lazygophers/log/constant"
	"github.com/lazygophers/log/hooks"
)

// TestHooksIntegration demonstrates all built-in hooks
func TestHooksIntegration(t *testing.T) {
	t.Run("BasicHook", func(t *testing.T) {
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
	})

	t.Run("SensitiveDataMask", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		maskHook := hooks.NewSensitiveDataMaskHook()
		logger.AddHook(maskHook)

		logger.Info("Email: test@example.com")
		logger.Info("Credit card: 4111-1111-1111-1111")
	})

	t.Run("LevelFilter", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		levelHook := hooks.NewLevelFilterHook(int(InfoLevel))
		logger.AddHook(levelHook)

		logger.Debug("debug message")
		logger.Info("info message")
	})

	t.Run("ContextEnrichment", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		fields := map[string]interface{}{
			"service": "test-service",
			"version": "1.0.0",
		}
		enrichHook := hooks.NewContextEnrichHook(fields)
		logger.AddHook(enrichHook)

		logger.Info("test message")
	})

	t.Run("PrefixAndSuffix", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		prefixHook := hooks.NewPrefixHook("[PREFIX] ")
		suffixHook := hooks.NewSuffixHook(" [SUFFIX]")
		logger.AddHooks(prefixHook, suffixHook)

		logger.Info("test message")
	})

	t.Run("MaxLength", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		maxHook := hooks.NewMaxLengthHook(10)
		logger.AddHook(maxHook)

		logger.Info("This is a very long message")
	})

	t.Run("MinLength", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		minHook := hooks.NewMinLengthHook(10)
		logger.AddHook(minHook)

		logger.Info("short")
		logger.Info("This is a long enough message")
	})

	t.Run("ConditionalHook", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		condition := func(entry interface{}) bool {
			if e, ok := entry.(*Entry); ok {
				return e.Level == ErrorLevel
			}
			return false
		}

		prefixHook := hooks.NewPrefixHook("[ERROR] ")
		conditionalHook := hooks.NewConditionalHook(condition, prefixHook)
		logger.AddHook(conditionalHook)

		logger.Info("info message")
		logger.Error("error message")
	})

	t.Run("ChainHook", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		chain := hooks.NewChainHook(
			hooks.NewPrefixHook("[CHAIN1] "),
			hooks.NewPrefixHook("[CHAIN2] "),
		)
		logger.AddHook(chain)

		logger.Info("test message")
	})

	t.Run("MessageFilter", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		filterHook := hooks.NewMessageFilterHook()
		filterHook.AddAllowPattern("important")
		logger.AddHook(filterHook)

		logger.Info("important message")
		logger.Info("other message")
	})

	t.Run("CustomMaskPattern", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		maskHook := hooks.NewSensitiveDataMaskHook()
		maskHook.AddPattern(`\b\d{3}-\d{3}-\d{4}\b`)
		logger.AddHook(maskHook)

		logger.Info("Call 555-123-4567")
	})

	t.Run("HookFiltering", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		hook := constant.HookFunc(func(entry interface{}) interface{} {
			if e, ok := entry.(*Entry); ok {
				if strings.Contains(e.Message, "skip") {
					return nil
				}
			}
			return entry
		})

		loggerWithHook := logger.AddHook(hook)

		loggerWithHook.Info("normal message")
		loggerWithHook.Info("skip this message")
	})

	t.Run("HookModification", func(t *testing.T) {
		logger := New()
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		hook := constant.HookFunc(func(entry interface{}) interface{} {
			if e, ok := entry.(*Entry); ok {
				e.Message = "[MODIFIED] " + e.Message
			}
			return entry
		})

		logger.AddHook(hook).Info("test")
	})

	t.Run("MultipleHooks", func(t *testing.T) {
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
	})

	t.Run("RemoveHooks", func(t *testing.T) {
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
	})
}
