package hooks

import (
	"strings"
	"testing"

	"github.com/lazygophers/log/constant"
)

func TestSensitiveDataMaskHook(t *testing.T) {
	t.Run("NewSensitiveDataMaskHook_creates_hook", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()
		if hook == nil {
			t.Error("NewSensitiveDataMaskHook should return non-nil")
		}
		if hook.mask != "***" {
			t.Errorf("Default mask should be ***, got %s", hook.mask)
		}
		if len(hook.patterns) == 0 {
			t.Error("Should have default patterns")
		}
		if len(hook.maskFields) == 0 {
			t.Error("Should have default mask fields")
		}
	})

	t.Run("SetMask_changes_mask", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()
		hook.SetMask("XXX")

		if hook.mask != "XXX" {
			t.Errorf("Mask should be XXX, got %s", hook.mask)
		}
	})

	t.Run("AddPattern_adds_custom_pattern", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()
		initialCount := len(hook.patterns)
		err := hook.AddPattern(`\b\d{3}\b`)
		if err != nil {
			t.Errorf("AddPattern should not error, got %v", err)
		}

		if len(hook.patterns) != initialCount+1 {
			t.Error("Pattern should be added")
		}
	})

	t.Run("AddPattern_with_invalid_regex", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()
		err := hook.AddPattern("[invalid(")
		if err == nil {
			t.Error("AddPattern should return error for invalid regex")
		}
	})

	t.Run("AddMaskField_adds_field", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()
		hook.AddMaskField("customField")

		if !hook.maskFields["customField"] {
			t.Error("Field should be added to maskFields")
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})

	t.Run("maskString_masks_patterns", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()

		// Test email masking
		result := hook.maskString("Contact test@example.com for support")
		if result == "Contact test@example.com for support" {
			t.Error("Email should be masked")
		}

		// Test credit card masking
		result = hook.maskString("Card: 4111-1111-1111-1111")
		if result == "Card: 4111-1111-1111-1111" {
			t.Error("Credit card should be masked")
		}

		// Test SSN masking
		result = hook.maskString("SSN: 123-45-6789")
		if result == "SSN: 123-45-6789" {
			t.Error("SSN should be masked")
		}
	})
}

func TestContextEnrichHook(t *testing.T) {
	t.Run("NewContextEnrichHook_with_fields", func(t *testing.T) {
		fields := map[string]interface{}{"service": "api", "version": "1.0"}
		hook := NewContextEnrichHook(fields)

		if hook == nil {
			t.Error("NewContextEnrichHook should return non-nil")
		}
		if hook.fields == nil {
			t.Error("Fields should be set")
		}
	})

	t.Run("NewContextEnrichHook_with_nil_fields", func(t *testing.T) {
		hook := NewContextEnrichHook(nil)

		if hook == nil {
			t.Error("NewContextEnrichHook should return non-nil")
		}
	})

	t.Run("AddField_adds_single_field", func(t *testing.T) {
		hook := NewContextEnrichHook(nil)
		hook.AddField("newField", "value")

		if hook.fields == nil {
			t.Error("Fields map should be created")
		}
		if hook.fields["newField"] != "value" {
			t.Error("Field should be added")
		}
	})

	t.Run("SetFields_replaces_fields", func(t *testing.T) {
		hook := NewContextEnrichHook(map[string]interface{}{"old": "value"})
		newFields := map[string]interface{}{"new": "value"}
		hook.SetFields(newFields)

		if hook.fields["old"] != nil {
			t.Error("Old fields should be replaced")
		}
		if hook.fields["new"] != "value" {
			t.Error("New fields should be set")
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewContextEnrichHook(map[string]interface{}{"key": "value"})
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewContextEnrichHook(map[string]interface{}{"key": "value"})
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})

	t.Run("OnWrite_with_empty_fields", func(t *testing.T) {
		hook := NewContextEnrichHook(nil)
		entry := struct{ Fields []interface{} }{Fields: []interface{}{}}

		result := hook.OnWrite(entry)
		if result == nil {
			t.Error("Should return entry when no fields to enrich")
		}
	})
}

func TestLevelFilterHook(t *testing.T) {
	t.Run("NewLevelFilterHook_creates_hook", func(t *testing.T) {
		hook := NewLevelFilterHook(5)
		if hook == nil {
			t.Error("NewLevelFilterHook should return non-nil")
		}
		if hook.minLevel != 5 {
			t.Errorf("MinLevel should be 5, got %d", hook.minLevel)
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewLevelFilterHook(5)
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewLevelFilterHook(5)
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})
}

func TestMessageFilterHook(t *testing.T) {
	t.Run("NewMessageFilterHook_creates_hook", func(t *testing.T) {
		hook := NewMessageFilterHook()
		if hook == nil {
			t.Error("NewMessageFilterHook should return non-nil")
		}
		// allow and deny start as nil, not empty slices
		if hook.allow != nil && len(hook.allow) != 0 {
			t.Error("Allow patterns should start empty")
		}
		if hook.deny != nil && len(hook.deny) != 0 {
			t.Error("Deny patterns should start empty")
		}
	})

	t.Run("AddAllowPattern_adds_pattern", func(t *testing.T) {
		hook := NewMessageFilterHook()
		err := hook.AddAllowPattern(`error`)
		if err != nil {
			t.Errorf("AddAllowPattern should not error, got %v", err)
		}
		if len(hook.allow) != 1 {
			t.Error("Pattern should be added to allow list")
		}
	})

	t.Run("AddDenyPattern_adds_pattern", func(t *testing.T) {
		hook := NewMessageFilterHook()
		err := hook.AddDenyPattern(`secret`)
		if err != nil {
			t.Errorf("AddDenyPattern should not error, got %v", err)
		}
		if len(hook.deny) != 1 {
			t.Error("Pattern should be added to deny list")
		}
	})

	t.Run("AddPattern_with_invalid_regex", func(t *testing.T) {
		hook := NewMessageFilterHook()
		err := hook.AddAllowPattern("[invalid")
		if err == nil {
			t.Error("Should return error for invalid regex")
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewMessageFilterHook()
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewMessageFilterHook()
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})
}

func TestFieldFilterHook(t *testing.T) {
	t.Run("NewFieldFilterHook_creates_hook", func(t *testing.T) {
		hook := NewFieldFilterHook()
		if hook == nil {
			t.Error("NewFieldFilterHook should return non-nil")
		}
		if hook.allowedFields == nil {
			t.Error("Allowed fields map should be initialized")
		}
		if hook.deniedFields == nil {
			t.Error("Denied fields map should be initialized")
		}
	})

	t.Run("AllowField_adds_allowed_value", func(t *testing.T) {
		hook := NewFieldFilterHook()
		hook.AllowField("level", "info")

		if hook.allowedFields["level"] == nil {
			t.Error("Field should be added to allowedFields")
		}
		if !hook.allowedFields["level"]["info"] {
			t.Error("Value should be marked as allowed")
		}
	})

	t.Run("DenyField_adds_denied_value", func(t *testing.T) {
		hook := NewFieldFilterHook()
		hook.DenyField("level", "debug")

		if hook.deniedFields["level"] == nil {
			t.Error("Field should be added to deniedFields")
		}
		if !hook.deniedFields["level"]["debug"] {
			t.Error("Value should be marked as denied")
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewFieldFilterHook()
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewFieldFilterHook()
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})
}

func TestMinLengthHook(t *testing.T) {
	t.Run("NewMinLengthHook_creates_hook", func(t *testing.T) {
		hook := NewMinLengthHook(10)
		if hook == nil {
			t.Error("NewMinLengthHook should return non-nil")
		}
		if hook.MinLength != 10 {
			t.Errorf("MinLength should be 10, got %d", hook.MinLength)
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewMinLengthHook(10)
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewMinLengthHook(10)
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})
}

func TestMaxLengthHook(t *testing.T) {
	t.Run("NewMaxLengthHook_creates_hook", func(t *testing.T) {
		hook := NewMaxLengthHook(10)
		if hook == nil {
			t.Error("NewMaxLengthHook should return non-nil")
		}
		if hook.MaxLength != 10 {
			t.Errorf("MaxLength should be 10, got %d", hook.MaxLength)
		}
		if hook.TruncateSuffix != "..." {
			t.Errorf("Default suffix should be ..., got %s", hook.TruncateSuffix)
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewMaxLengthHook(10)
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewMaxLengthHook(10)
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})
}

func TestPrefixHook(t *testing.T) {
	t.Run("NewPrefixHook_creates_hook", func(t *testing.T) {
		hook := NewPrefixHook("[PREFIX] ")
		if hook == nil {
			t.Error("NewPrefixHook should return non-nil")
		}
		if hook.Prefix != "[PREFIX] " {
			t.Errorf("Prefix should be [PREFIX], got %s", hook.Prefix)
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewPrefixHook("[TEST] ")
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewPrefixHook("[TEST] ")
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})
}

func TestSuffixHook(t *testing.T) {
	t.Run("NewSuffixHook_creates_hook", func(t *testing.T) {
		hook := NewSuffixHook(" [SUFFIX]")
		if hook == nil {
			t.Error("NewSuffixHook should return non-nil")
		}
		if hook.Suffix != " [SUFFIX]" {
			t.Errorf("Suffix should be [SUFFIX], got %s", hook.Suffix)
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewSuffixHook(" [TEST]")
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_handles_non_entry_type", func(t *testing.T) {
		hook := NewSuffixHook(" [TEST]")
		result := hook.OnWrite("string")
		if result != "string" {
			t.Error("Should return original entry for non-entryLike type")
		}
	})
}

func TestConditionalHook(t *testing.T) {
	t.Run("NewConditionalHook_creates_hook", func(t *testing.T) {
		condition := func(entry interface{}) bool { return true }
		hook := NewConditionalHook(condition, nil)
		if hook == nil {
			t.Error("NewConditionalHook should return non-nil")
		}
		if hook.condition == nil {
			t.Error("Condition should be set")
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		condition := func(entry interface{}) bool { return true }
		hook := NewConditionalHook(condition, nil)
		result := hook.OnWrite(nil)
		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_skips_hook_when_condition_false", func(t *testing.T) {
		executed := false
		innerHook := constant.HookFunc(func(entry interface{}) interface{} {
			executed = true
			return entry
		})
		condition := func(entry interface{}) bool { return false }

		hook := NewConditionalHook(condition, innerHook)
		result := hook.OnWrite("test")

		if executed {
			t.Error("Inner hook should not be executed when condition is false")
		}
		if result != "test" {
			t.Error("Should return original entry when condition is false")
		}
	})

	t.Run("OnWrite_executes_hook_when_condition_true", func(t *testing.T) {
		executed := false
		innerHook := constant.HookFunc(func(entry interface{}) interface{} {
			executed = true
			return entry
		})
		condition := func(entry interface{}) bool { return true }

		hook := NewConditionalHook(condition, innerHook)
		result := hook.OnWrite("test")

		if !executed {
			t.Error("Inner hook should be executed when condition is true")
		}
		if result != "test" {
			t.Error("Should return result from inner hook")
		}
	})
}

func TestChainHook(t *testing.T) {
	t.Run("NewChainHook_creates_hook", func(t *testing.T) {
		hook := NewChainHook()
		if hook == nil {
			t.Error("NewChainHook should return non-nil")
		}
		// hooks starts as nil when no hooks passed
		if hook.hooks != nil && len(hook.hooks) != 0 {
			t.Error("Hooks should be empty when none passed")
		}
	})

	t.Run("NewChainHook_with_hooks", func(t *testing.T) {
		hook1 := constant.HookFunc(func(entry interface{}) interface{} {
			return entry
		})
		hook2 := constant.HookFunc(func(entry interface{}) interface{} {
			return entry
		})

		hook := NewChainHook(hook1, hook2)
		if len(hook.hooks) != 2 {
			t.Error("Hooks should be stored")
		}
	})

	t.Run("OnWrite_handles_nil_entry", func(t *testing.T) {
		hook := NewChainHook()
		result := hook.OnWrite(nil)

		if result != nil {
			t.Error("Should return nil for nil entry")
		}
	})

	t.Run("OnWrite_with_empty_hooks", func(t *testing.T) {
		hook := NewChainHook()
		entry := "test"
		result := hook.OnWrite(entry)

		if result != entry {
			t.Error("Should return entry unchanged for empty hooks")
		}
	})

	t.Run("OnWrite_executes_hooks_in_sequence", func(t *testing.T) {
		order := []string{}
		hook1 := constant.HookFunc(func(entry interface{}) interface{} {
			order = append(order, "hook1")
			return entry
		})
		hook2 := constant.HookFunc(func(entry interface{}) interface{} {
			order = append(order, "hook2")
			return entry
		})

		hook := NewChainHook(hook1, hook2)
		result := hook.OnWrite("test")

		if len(order) != 2 || order[0] != "hook1" || order[1] != "hook2" {
			t.Error("Hooks should execute in sequence")
		}
		if result != "test" {
			t.Error("Should return final result")
		}
	})

	t.Run("OnWrite_stops_on_nil_filter", func(t *testing.T) {
		hook1 := constant.HookFunc(func(entry interface{}) interface{} {
			return nil // Filter out
		})
		called := false
		hook2 := constant.HookFunc(func(entry interface{}) interface{} {
			called = true
			return entry
		})

		hook := NewChainHook(hook1, hook2)
		result := hook.OnWrite("test")

		if result != nil {
			t.Error("Should return nil when hook filters")
		}
		if called {
			t.Error("Should stop chain on nil result")
		}
	})
}

// --- Integration tests with real *constant.Entry ---

func TestSensitiveDataMaskHook_WithEntry(t *testing.T) {
	hook := NewSensitiveDataMaskHook()
	entry := &constant.Entry{
		Message: "user admin@test.com logged in, card 4111-1111-1111-1111",
		Fields: []constant.KV{
			{Key: "password", Value: "secret123"},
			{Key: "token", Value: "abc123"},
			{Key: "name", Value: "John"},
			{Key: "contact", Value: "john@example.com"},
			{Key: "age", Value: 30},
		},
		File:       "handler.go",
		CallerName: "app.handle",
	}

	result := hook.OnWrite(entry)
	e := result.(*constant.Entry)

	if strings.Contains(e.Message, "admin@test.com") {
		t.Error("email in message should be masked")
	}
	if strings.Contains(e.Message, "4111-1111-1111-1111") {
		t.Error("card number in message should be masked")
	}

	for _, f := range e.Fields {
		switch f.Key {
		case "password", "token":
			if f.Value != "***" {
				t.Errorf("%s field should be masked by key, got %v", f.Key, f.Value)
			}
		case "name":
			if f.Value != "John" {
				t.Error("name field should not be masked")
			}
		case "contact":
			if s, ok := f.Value.(string); ok && strings.Contains(s, "john@example.com") {
				t.Error("contact email value should be masked")
			}
		case "age":
			if f.Value != 30 {
				t.Error("age field should not be masked")
			}
		}
	}
}

func TestContextEnrichHook_WithEntry(t *testing.T) {
	hook := NewContextEnrichHook(map[string]interface{}{"service": "api", "version": "1.0"})
	entry := &constant.Entry{Message: "test"}

	result := hook.OnWrite(entry)
	e := result.(*constant.Entry)

	if len(e.Fields) != 2 {
		t.Fatalf("expected 2 enriched fields, got %d", len(e.Fields))
	}

	found := map[string]bool{}
	for _, f := range e.Fields {
		found[f.Key] = true
	}
	if !found["service"] || !found["version"] {
		t.Error("both service and version fields should be enriched")
	}
}

func TestLevelFilterHook_WithEntry(t *testing.T) {
	// constant.Level: Panic=0, Fatal=1, Error=2, Warn=3, Info=4, Debug=5, Trace=6
	// minLevel=4 means: filter entries where int(Level) < 4 (i.e. Error, Fatal, Panic)
	hook := NewLevelFilterHook(4)

	filteredEntry := &constant.Entry{Level: constant.Level(2)} // Error: 2 < 4 → filtered
	result := hook.OnWrite(filteredEntry)
	if result != nil {
		t.Error("error level entry should be filtered when minLevel=4")
	}

	passEntry := &constant.Entry{Level: constant.Level(4)} // Info: 4 < 4 → false → passes
	result = hook.OnWrite(passEntry)
	if result == nil {
		t.Error("info level entry should pass when minLevel=4")
	}

	traceEntry := &constant.Entry{Level: constant.Level(6)} // Trace: 6 < 4 → false → passes
	result = hook.OnWrite(traceEntry)
	if result == nil {
		t.Error("trace level entry should pass when minLevel=4")
	}
}

func TestMessageFilterHook_WithEntry(t *testing.T) {
	t.Run("deny_pattern", func(t *testing.T) {
		hook := NewMessageFilterHook()
		hook.AddDenyPattern(`secret`)

		denied := &constant.Entry{Message: "this is a secret message"}
		result := hook.OnWrite(denied)
		if result != nil {
			t.Error("message matching deny pattern should be filtered")
		}
	})

	t.Run("allow_deny_combined", func(t *testing.T) {
		hook := NewMessageFilterHook()
		hook.AddAllowPattern(`error`)
		hook.AddDenyPattern(`internal`)

		// Denied takes precedence
		result := hook.OnWrite(&constant.Entry{Message: "internal error occurred"})
		if result != nil {
			t.Error("denied pattern should take precedence")
		}

		// Allowed
		result = hook.OnWrite(&constant.Entry{Message: "error in module"})
		if result == nil {
			t.Error("allowed message should pass")
		}

		// No allow match
		result = hook.OnWrite(&constant.Entry{Message: "info message"})
		if result != nil {
			t.Error("message not matching allow should be filtered")
		}
	})
}

func TestFieldFilterHook_WithEntry(t *testing.T) {
	hook := NewFieldFilterHook()
	hook.AllowField("env", "production")
	hook.DenyField("user", "admin")

	denied := &constant.Entry{
		Fields: []constant.KV{{Key: "user", Value: "admin"}},
	}
	if result := hook.OnWrite(denied); result != nil {
		t.Error("denied field value should be filtered")
	}

	allowed := &constant.Entry{
		Fields: []constant.KV{{Key: "env", Value: "production"}},
	}
	if result := hook.OnWrite(allowed); result == nil {
		t.Error("allowed field value should pass")
	}

	notAllowed := &constant.Entry{
		Fields: []constant.KV{{Key: "env", Value: "staging"}},
	}
	if result := hook.OnWrite(notAllowed); result != nil {
		t.Error("non-allowed field value should be filtered")
	}
}

func TestMinLengthHook_WithEntry(t *testing.T) {
	hook := NewMinLengthHook(5)

	if result := hook.OnWrite(&constant.Entry{Message: "hi"}); result != nil {
		t.Error("short message should be filtered")
	}
	if result := hook.OnWrite(&constant.Entry{Message: "hello world"}); result == nil {
		t.Error("long enough message should pass")
	}
}

func TestMaxLengthHook_WithEntry(t *testing.T) {
	hook := NewMaxLengthHook(5)

	long := &constant.Entry{Message: "hello world"}
	result := hook.OnWrite(long)
	e := result.(*constant.Entry)
	if !strings.HasSuffix(e.Message, "...") {
		t.Error("long message should be truncated with suffix")
	}
	if len([]rune(e.Message)) != 5+3 {
		t.Errorf("expected 8 runes, got %d: %q", len([]rune(e.Message)), e.Message)
	}

	short := &constant.Entry{Message: "hi"}
	result = hook.OnWrite(short)
	e = result.(*constant.Entry)
	if e.Message != "hi" {
		t.Error("short message should not be truncated")
	}
}

func TestPrefixHook_WithEntry(t *testing.T) {
	hook := NewPrefixHook("[APP] ")
	entry := &constant.Entry{Message: "started"}
	e := hook.OnWrite(entry).(*constant.Entry)
	if e.Message != "[APP] started" {
		t.Errorf("expected prefix prepended, got %q", e.Message)
	}
}

func TestSuffixHook_WithEntry(t *testing.T) {
	hook := NewSuffixHook(" [END]")
	entry := &constant.Entry{Message: "done"}
	e := hook.OnWrite(entry).(*constant.Entry)
	if e.Message != "done [END]" {
		t.Errorf("expected suffix appended, got %q", e.Message)
	}
}

func TestConditionalHook_WithEntry(t *testing.T) {
	inner := NewPrefixHook("[ERR]")
	hook := NewConditionalHook(
		func(entry interface{}) bool {
			if e, ok := entry.(*constant.Entry); ok {
				return int(e.Level) <= 2 // Error(2), Fatal(1), Panic(0)
			}
			return false
		},
		inner,
	)

	e := hook.OnWrite(&constant.Entry{Level: constant.Level(2), Message: "error"}).(*constant.Entry)
	if e.Message != "[ERR]error" {
		t.Errorf("error entry should have prefix, got %q", e.Message)
	}

	e = hook.OnWrite(&constant.Entry{Level: constant.Level(4), Message: "info"}).(*constant.Entry)
	if e.Message != "info" {
		t.Errorf("info entry should be unchanged, got %q", e.Message)
	}
}

func TestChainHook_WithEntry(t *testing.T) {
	t.Run("chain_prefix_suffix", func(t *testing.T) {
		chain := NewChainHook(NewPrefixHook("[A]"), NewSuffixHook("[B]"))
		e := chain.OnWrite(&constant.Entry{Message: "msg"}).(*constant.Entry)
		if e.Message != "[A]msg[B]" {
			t.Errorf("expected [A]msg[B], got %q", e.Message)
		}
	})

	t.Run("chain_with_filter", func(t *testing.T) {
		chain := NewChainHook(NewMinLengthHook(10), NewPrefixHook("[A]"))
		if result := chain.OnWrite(&constant.Entry{Message: "hi"}); result != nil {
			t.Error("chain should stop after filter returns nil")
		}
	})
}
