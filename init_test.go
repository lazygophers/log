package log

import (
	"testing"
)

func TestPid(t *testing.T) {
	pid := Pid()
	if pid <= 0 {
		t.Errorf("Expected positive PID, got %d", pid)
	}
}

func TestNew(t *testing.T) {
	logger := New()
	if logger == nil {
		t.Fatal("New() returned nil")
	}
	if logger.level != DebugLevel {
		t.Errorf("Expected default level DebugLevel, got %v", logger.level)
	}
}

func TestClone(t *testing.T) {
	cloned := Clone()
	if cloned == nil {
		t.Fatal("Clone() returned nil")
	}
	if cloned.level != std.level {
		t.Errorf("Cloned logger level mismatch: expected %v, got %v", std.level, cloned.level)
	}
}

func TestSync(t *testing.T) {
	// Just verify it doesn't panic
	Sync()
}

func TestSetCallerDepth(t *testing.T) {
	result := SetCallerDepth(5)
	if result != std {
		t.Error("SetCallerDepth should return std logger")
	}
	if std.callerDepth != 5 {
		t.Errorf("Expected callerDepth 5, got %d", std.callerDepth)
	}
	// Reset to default
	std.callerDepth = 4
}

func TestSetPrefixMsg(t *testing.T) {
	result := SetPrefixMsg("[test]")
	if result != std {
		t.Error("SetPrefixMsg should return std logger")
	}
	if string(std.PrefixMsg) != "[test]" {
		t.Errorf("Expected PrefixMsg [test], got %s", string(std.PrefixMsg))
	}
	// Reset
	std.PrefixMsg = nil
}

func TestAppendPrefixMsg(t *testing.T) {
	SetPrefixMsg("[a]")
	result := AppendPrefixMsg("[b]")
	if result != std {
		t.Error("AppendPrefixMsg should return std logger")
	}
	if string(std.PrefixMsg) != "[a][b]" {
		t.Errorf("Expected PrefixMsg [a][b], got %s", string(std.PrefixMsg))
	}
	// Reset
	std.PrefixMsg = nil
}

func TestSetSuffixMsg(t *testing.T) {
	result := SetSuffixMsg("[test]")
	if result != std {
		t.Error("SetSuffixMsg should return std logger")
	}
	if string(std.SuffixMsg) != "[test]" {
		t.Errorf("Expected SuffixMsg [test], got %s", string(std.SuffixMsg))
	}
	// Reset
	std.SuffixMsg = nil
}

func TestAppendSuffixMsg(t *testing.T) {
	SetSuffixMsg("[a]")
	result := AppendSuffixMsg("[b]")
	if result != std {
		t.Error("AppendSuffixMsg should return std logger")
	}
	if string(std.SuffixMsg) != "[a][b]" {
		t.Errorf("Expected SuffixMsg [a][b], got %s", string(std.SuffixMsg))
	}
	// Reset
	std.SuffixMsg = nil
}

func TestParsingAndEscaping(t *testing.T) {
	result := ParsingAndEscaping(true)
	if result != std {
		t.Error("ParsingAndEscaping should return std logger")
	}
}

func TestCaller(t *testing.T) {
	result := Caller(true)
	if result != std {
		t.Error("Caller should return std logger")
	}
}
