package log

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatter_Format_WithPrefixAndSuffix(t *testing.T) {
	entry := &Entry{
		PrefixMsg:  []byte("PREFIX"),
		Message:    "test message\nwith newline",
		SuffixMsg:  []byte("SUFFIX"),
		Level:      DebugLevel,
		CallerDir:  "github.com/lazygophers/log",
		File:       "formatter.go",
		CallerLine: 0,
		CallerFunc: "main",
		Pid:        1,
		Gid:        456,
	}

	formatter := &Formatter{
		DisableParsingAndEscaping: false,
		DisableCaller:             false,
	}

	result := formatter.Format(entry)

	assert.Contains(t, string(result), "PREFIX")
	assert.Contains(t, string(result), "SUFFIX", "Suffix missing")
	assert.Contains(t, string(result), "\u001B[32m", "确保显示调试颜色代码")
	assert.Contains(t, string(result), "\n", "允许换行符存在")
}

func TestFormatter_Format_DisableParsing(t *testing.T) {
	entry := &Entry{
		Message: "test message with\nnewline",
		Level:   ErrorLevel,
	}

	formatter := &Formatter{
		DisableParsingAndEscaping: true,
	}

	result := formatter.Format(entry)
	assert.Contains(t, string(result), "\n", "允许换行符存在")
}

func BenchmarkFormatter_Format(b *testing.B) {
	entry := &Entry{
		Message:    "benchmark message",
		Level:      InfoLevel,
		CallerDir:  "github.com/lazygophers/log",
		File:       "formatter.go",
		CallerLine: 0,
		CallerFunc: "main",
		Pid:        1,
		Gid:        456,
	}

	formatter := &Formatter{
		DisableParsingAndEscaping: false,
		DisableCaller:             false,
	}

	for i := 0; i < b.N; i++ {
		formatter.Format(entry)
	}
}

func TestColorByLevel(t *testing.T) {
	tests := []struct {
		level    Level
		expected []byte
	}{
		{DebugLevel, colorGreen},
		{TraceLevel, colorGreen},
		{WarnLevel, colorYellow},
		{ErrorLevel, colorRed},
		{FatalLevel, colorRed},
		{PanicLevel, colorRed},
		{InfoLevel, colorGreen},
	}

	for _, test := range tests {
		if color := getColorByLevel(test.level); string(color) != string(test.expected) {
			t.Errorf("getColorByLevel(%v) = %s, expected %s",
				test.level, string(color), string(test.expected))
		}
	}
}

func TestFormatter_Clone(t *testing.T) {
	original := &Formatter{
		Module:                    "test-module",
		DisableParsingAndEscaping: true,
		DisableCaller:             false,
	}

	cloned := original.Clone()
	if !reflect.DeepEqual(original, cloned) || reflect.ValueOf(original.Format).Pointer() == reflect.ValueOf(cloned.Format).Pointer() {
		t.Errorf("Clone() produced same formatter: %v vs %v", original, cloned)
	}
}
