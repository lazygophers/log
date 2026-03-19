package log

import (
	"errors"
	"fmt"
	"testing"
)

func TestFastStringify_AllTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"string", "hello", "hello"},
		{"empty_string", "", ""},
		{"bytes", []byte("hello"), "hello"},
		{"empty_bytes", []byte(""), ""},
		{"int", 42, "42"},
		{"int_zero", 0, "0"},
		{"int_negative", -42, "-42"},
		{"int8", int8(42), "42"},
		{"int8_min", int8(-128), "-128"},
		{"int8_max", int8(127), "127"},
		{"int16", int16(42), "42"},
		{"int16_min", int16(-32768), "-32768"},
		{"int16_max", int16(32767), "32767"},
		{"int32", int32(42), "42"},
		{"int32_min", int32(-2147483648), "-2147483648"},
		{"int32_max", int32(2147483647), "2147483647"},
		{"int64", int64(42), "42"},
		{"int64_large", int64(9223372036854775807), "9223372036854775807"},
		{"uint", uint(42), "42"},
		{"uint_zero", uint(0), "0"},
		{"uint8", uint8(42), "42"},
		{"uint8_max", uint8(255), "255"},
		{"uint16", uint16(42), "42"},
		{"uint16_max", uint16(65535), "65535"},
		{"uint32", uint32(42), "42"},
		{"uint32_max", uint32(4294967295), "4294967295"},
		{"uint64", uint64(42), "42"},
		{"uint64_max", uint64(18446744073709551615), "18446744073709551615"},
		{"float32", float32(3.14), "3.14"},
		{"float32_zero", float32(0), "0"},
		{"float32_negative", float32(-1.5), "-1.5"},
		{"float64", 3.14159, "3.14159"},
		{"float64_zero", float64(0), "0"},
		{"float64_negative", -2.718, "-2.718"},
		{"bool_true", true, "true"},
		{"bool_false", false, "false"},
		{"error", errors.New("test error"), "test error"},
		{"nil_error", (error)(nil), ""},
		{"nil", nil, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fastStringify(tt.input)
			if result != tt.expected {
				t.Errorf("fastStringify(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFastStringify_FallbackType(t *testing.T) {
	// struct type should fallback to fmt.Sprint
	type custom struct{ X int }
	input := custom{X: 42}
	result := fastStringify(input)
	expected := fmt.Sprint(input)
	if result != expected {
		t.Errorf("fastStringify(%v) = %q, want %q", input, result, expected)
	}
}

func TestFastStringify_NilErrorInterface(t *testing.T) {
	// explicitly typed nil error
	var err error
	result := fastStringify(err)
	if result != "" {
		t.Errorf("fastStringify(nil error) = %q, want empty string", result)
	}
}

func TestFastSprint_ZeroArgs(t *testing.T) {
	result := fastSprint()
	if result != "" {
		t.Errorf("fastSprint() = %q, want empty string", result)
	}
}

func TestFastSprint_SingleArg(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"string", "hello", "hello"},
		{"int", 42, "42"},
		{"bool", true, "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fastSprint(tt.input)
			if result != tt.expected {
				t.Errorf("fastSprint(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFastSprint_MultipleArgs(t *testing.T) {
	result := fastSprint("hello", " ", "world")
	expected := fmt.Sprint("hello", " ", "world")
	if result != expected {
		t.Errorf("fastSprint multiple args = %q, want %q", result, expected)
	}
}

func TestFastSprintf_NoArgs(t *testing.T) {
	result := fastSprintf("hello world")
	if result != "hello world" {
		t.Errorf("fastSprintf no args = %q, want %q", result, "hello world")
	}
}

func TestFastSprintf_SinglePercentS(t *testing.T) {
	result := fastSprintf("%s", "hello")
	if result != "hello" {
		t.Errorf("fastSprintf %%s = %q, want %q", result, "hello")
	}
}

func TestFastSprintf_StandardFormat(t *testing.T) {
	result := fastSprintf("value: %d", 42)
	expected := "value: 42"
	if result != expected {
		t.Errorf("fastSprintf standard = %q, want %q", result, expected)
	}
}

func TestFastSprintf_MultipleArgs(t *testing.T) {
	result := fastSprintf("%s=%d", "key", 42)
	expected := fmt.Sprintf("%s=%d", "key", 42)
	if result != expected {
		t.Errorf("fastSprintf multiple = %q, want %q", result, expected)
	}
}
