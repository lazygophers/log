package log

import (
	"strings"
	"testing"
	"time"
)

// TestJSONFormatterEscapeStringCoverage 测试 jsonEscapeString 的所有分支
func TestJSONFormatterEscapeStringCoverage(t *testing.T) {
	t.Run("jsonEscapeString_all_branches", func(t *testing.T) {
		testCases := []struct {
			name  string
			input string
			// jsonEscapeString 的输出
			// 例如: "test\"" -> "test\\\""
			// 在 Go 字符串字面量中:
			// - 输入的 \" 是一个引号字符
			// - 输出的 \\\" 是三个字符: \, \, "
		}{
			{
				name:  "quote",
				input: "test\"quote", // 输入: 引号字符
			},
			{
				name:  "backslash",
				input: "test\\backslash", // 输入: 反斜杠字符
			},
			{
				name:  "newline",
				input: "test\nnewline", // 输入: 换行符
			},
			{
				name:  "carriage_return",
				input: "test\rreturn", // 输入: 回车符
			},
			{
				name:  "tab",
				input: "test\ttab", // 输入: 制表符
			},
			{
				name:  "control_chars_0x00_to_0x0f",
				input: "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x0b\x0c\x0e\x0f",
			},
			{
				name:  "control_chars_0x10_to_0x1f",
				input: "\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f",
			},
			{
				name:  "unicode_above_32",
				input: "ABC中文123",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// 调用 jsonEscapeString 确保不会 panic
				result := jsonEscapeString(tc.input)

				// 基本检查：结果不应该为空
				if result == "" && tc.input != "" {
					t.Errorf("jsonEscapeString(%q) should return non-empty", tc.input)
				}

				// 验证控制字符被转义为 \u00XX 格式
				if tc.input < "\x20" {
					if !strings.Contains(result, `\u00`) {
						t.Errorf("jsonEscapeString(%q) should contain \\u00XX, got %q", tc.input, result)
					}
				}
			})
		}
	})

	t.Run("hexByte_all_values", func(t *testing.T) {
		// 测试 hexByte 的所有可能值（0-15）
		// 0-9 应该使用 '0' + b 分支
		// 10-15 应该使用 'a' + b - 10 分支

		for b := byte(0); b < 16; b++ {
			result := hexByte(b)

			var expected byte
			if b < 10 {
				expected = '0' + b
			} else {
				expected = 'a' + b - 10
			}

			if result != expected {
				t.Errorf("hexByte(%d) = %c, want %c", b, result, expected)
			}
		}
	})
}

// TestJSONFormatterMarshallError 触发 JSON 序列化错误路径
func TestJSONFormatterMarshallError(t *testing.T) {
	t.Run("channel_field_triggers_error_path", func(t *testing.T) {
		formatter := &JSONFormatter{}

		ch := make(chan int)
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test with \"quotes\" and \\backslash",
			Time:    time.Now(),
			Pid:     123,
			Fields: []KV{
				{Key: "channel", Value: ch},
			},
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		// 应该触发错误路径并使用 jsonEscapeString
		if !strings.Contains(resultStr, "JSON marshaling failed") {
			t.Error("Should include marshaling error message")
		}

		// 验证 jsonEscapeString 被调用（原始消息被包含）
		if !strings.Contains(resultStr, "test with") {
			t.Error("Should include original message")
		}
	})
}
