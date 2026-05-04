package hooks

import (
	"testing"
)

// TestHooksCoverage 提升hooks包的覆盖率
func TestHooksCoverage(t *testing.T) {
	t.Run("SensitiveDataMaskHook_actual_entry_processing", func(t *testing.T) {
		hook := NewSensitiveDataMaskHook()

		// 创建匹配entryLike的匿名结构体
		type testEntry struct {
			Message    string
			Fields     []interface{}
			File       string
			CallerName string
		}

		entry := testEntry{
			Message:    "Contact test@example.com for support, card: 4111-1111-1111-1111",
			Fields:     []interface{}{map[string]interface{}{"password": "secret123", "user": "admin"}},
			File:       "/path/to/file.go",
			CallerName: "main.Func",
		}

		result := hook.OnWrite(entry)

		// 验证结果
		if result == nil {
			t.Fatal("OnWrite should return entry")
		}

		// 类型断言检查结果
		if r, ok := result.(testEntry); ok {
			// 检查message中的敏感数据被mask
			if r.Message == entry.Message {
				t.Error("Message should be masked")
			}
			if r.File == entry.File {
				t.Error("File should be masked")
			}
		}
	})

	t.Run("ContextEnrichHook_adds_fields_to_entry", func(t *testing.T) {
		fields := map[string]interface{}{
			"service": "api",
			"version": "1.0",
		}
		hook := NewContextEnrichHook(fields)

		type testEntry struct {
			Fields []interface{}
		}

		entry := testEntry{
			Fields: []interface{}{
				map[string]interface{}{"existing": "field"},
			},
		}

		result := hook.OnWrite(entry)

		if result == nil {
			t.Fatal("OnWrite should return entry")
		}

		if r, ok := result.(testEntry); ok {
			if len(r.Fields) != 3 { // 1 original + 2 added
				t.Errorf("Fields should have 3 entries, got %d", len(r.Fields))
			}
		}
	})

	t.Run("LevelFilterHook_filters_by_level", func(t *testing.T) {
		hook := NewLevelFilterHook(5) // 只允许级别 >= 5 的日志

		type testEntry struct {
			Level int
		}

		// 测试级别不够的情况
		entry := testEntry{Level: 3}
		result := hook.OnWrite(entry)
		if result != nil {
			t.Error("Should filter out entry with level < minLevel")
		}

		// 测试级别足够的情况
		entry = testEntry{Level: 10}
		result = hook.OnWrite(entry)
		if result == nil {
			t.Error("Should allow entry with level >= minLevel")
		}
	})

	t.Run("MessageFilterHook_allow_patterns", func(t *testing.T) {
		hook := NewMessageFilterHook()
		hook.AddAllowPattern(`error`)
		hook.AddAllowPattern(`warning`)

		type testEntry struct {
			Message string
		}

		// 匹配allow模式
		entry := testEntry{Message: "This is an error message"}
		result := hook.OnWrite(entry)
		if result == nil {
			t.Error("Should allow message matching allow pattern")
		}

		// 不匹配allow模式
		entry = testEntry{Message: "This is info message"}
		result = hook.OnWrite(entry)
		if result != nil {
			t.Error("Should filter out message not matching allow pattern")
		}
	})

	t.Run("MessageFilterHook_deny_patterns", func(t *testing.T) {
		hook := NewMessageFilterHook()
		hook.AddDenyPattern(`secret`)

		type testEntry struct {
			Message string
		}

		// 匹配deny模式
		entry := testEntry{Message: "This message contains secret data"}
		result := hook.OnWrite(entry)
		if result != nil {
			t.Error("Should filter out message matching deny pattern")
		}

		// 不匹配deny模式
		entry = testEntry{Message: "This is normal message"}
		result = hook.OnWrite(entry)
		if result == nil {
			t.Error("Should allow message not matching deny pattern")
		}
	})

	t.Run("FieldFilterHook_filters_by_field_values", func(t *testing.T) {
		hook := NewFieldFilterHook()
		hook.AllowField("level", "info")
		hook.AllowField("level", "debug")

		type testEntry struct {
			Fields []interface{}
		}

		// 允许的字段值
		entry := testEntry{
			Fields: []interface{}{
				map[string]interface{}{"level": "info"},
			},
		}
		result := hook.OnWrite(entry)
		if result == nil {
			t.Error("Should allow entry with allowed field value")
		}

		// 不允许的字段值
		entry = testEntry{
			Fields: []interface{}{
				map[string]interface{}{"level": "error"},
			},
		}
		result = hook.OnWrite(entry)
		if result != nil {
			t.Error("Should filter out entry with disallowed field value")
		}
	})

	t.Run("FieldFilterHook_denies_field_values", func(t *testing.T) {
		hook := NewFieldFilterHook()
		hook.DenyField("level", "debug")

		type testEntry struct {
			Fields []interface{}
		}

		// 被拒绝的字段值
		entry := testEntry{
			Fields: []interface{}{
				map[string]interface{}{"level": "debug"},
			},
		}
		result := hook.OnWrite(entry)
		if result != nil {
			t.Error("Should filter out entry with denied field value")
		}
	})

	t.Run("MinLengthHook_filters_short_messages", func(t *testing.T) {
		hook := NewMinLengthHook(10)

		type testEntry struct {
			Message string
		}

		// 短消息
		entry := testEntry{Message: "short"}
		result := hook.OnWrite(entry)
		if result != nil {
			t.Error("Should filter out short message")
		}

		// 足够长的消息
		entry = testEntry{Message: "This is a long enough message"}
		result = hook.OnWrite(entry)
		if result == nil {
			t.Error("Should allow long message")
		}
	})

	t.Run("MaxLengthHook_truncates_long_messages", func(t *testing.T) {
		hook := NewMaxLengthHook(10)

		type testEntry struct {
			Message string
		}

		entry := testEntry{Message: "This is a very long message that should be truncated"}
		result := hook.OnWrite(entry)

		if result == nil {
			t.Fatal("OnWrite should return entry")
		}

		if r, ok := result.(testEntry); ok {
			if len(r.Message) > 13 { // 10 + "..."
				t.Errorf("Message should be truncated, got length %d", len(r.Message))
			}
		}
	})

	t.Run("PrefixHook_adds_prefix", func(t *testing.T) {
		hook := NewPrefixHook("[PREFIX] ")

		type testEntry struct {
			Message string
		}

		entry := testEntry{Message: "test message"}
		result := hook.OnWrite(entry)

		if r, ok := result.(testEntry); ok {
			if r.Message != "[PREFIX] test message" {
				t.Errorf("Message should have prefix, got %s", r.Message)
			}
		}
	})

	t.Run("SuffixHook_adds_suffix", func(t *testing.T) {
		hook := NewSuffixHook(" [SUFFIX]")

		type testEntry struct {
			Message string
		}

		entry := testEntry{Message: "test message"}
		result := hook.OnWrite(entry)

		if r, ok := result.(testEntry); ok {
			if r.Message != "test message [SUFFIX]" {
				t.Errorf("Message should have suffix, got %s", r.Message)
			}
		}
	})
}
