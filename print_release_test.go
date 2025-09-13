//go:build release && !discard

package log

import (
	"testing"
)

// TestDebugRelease tests the Debug function in release mode
// In release mode, Debug should be a no-op function
func TestDebugRelease(t *testing.T) {
	// Test that Debug function can be called without issues
	Debug("test debug message")
	Debug("another message", "with", "multiple", "args")
	Debug(123, true, nil)
	
	// Call it multiple times to ensure coverage
	for i := 0; i < 3; i++ {
		Debug("iteration", i)
	}
	
	// Since it's a no-op, we just ensure no panic occurs
	// and the function completes successfully
}

// TestDebugfRelease tests the Debugf function in release mode
// In release mode, Debugf should be a no-op function
func TestDebugfRelease(t *testing.T) {
	// Test that Debugf function can be called without issues
	Debugf("test debug message with format: %s", "value")
	Debugf("number: %d, bool: %v", 42, true)
	Debugf("empty format")
	
	// Call it multiple times to ensure coverage
	for i := 0; i < 3; i++ {
		Debugf("iteration %d", i)
	}
	
	// Since it's a no-op, we just ensure no panic occurs
	// and the function completes successfully
}

// TestReleaseLogPath tests the ReleaseLogPath variable
func TestReleaseLogPath(t *testing.T) {
	// Verify that ReleaseLogPath is set to a reasonable default
	if ReleaseLogPath == "" {
		t.Error("ReleaseLogPath should not be empty")
	}
	
	// Verify it contains expected path components
	if !contains(ReleaseLogPath, "lazygophers") {
		t.Error("ReleaseLogPath should contain 'lazygophers'")
	}
	if !contains(ReleaseLogPath, "log") {
		t.Error("ReleaseLogPath should contain 'log'")
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && 
		(s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || 
		s[len(s)-len(substr):] == substr || 
		containsSubstring(s, substr)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}