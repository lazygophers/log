package log

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"testing"
)

// envTestResult is used to pass results from subprocess back to parent via stdout.
type envTestResult struct {
	Level int `json:"level"`
}

// runEnvSubprocess runs the current test binary in a subprocess with given APP_ENV,
// and returns the log level that was set after init().
func runEnvSubprocess(t *testing.T, testName, appEnv string) Level {
	t.Helper()
	cmd := exec.Command(os.Args[0], "-test.run=^"+testName+"$")
	cmd.Env = append(os.Environ(), "TEST_ENV_INIT=1", "APP_ENV="+appEnv)
	cmd.Stderr = io.Discard
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("subprocess failed: %v", err)
	}
	var result envTestResult
	if err := json.Unmarshal(out, &result); err != nil {
		t.Fatalf("failed to parse subprocess output: %v, raw: %s", err, string(out))
	}
	return Level(result.Level)
}

func TestEnvInit_DevEnvironment(t *testing.T) {
	if os.Getenv("TEST_ENV_INIT") == "1" {
		// In subprocess: init() has already run with APP_ENV set.
		// Output the current level as JSON.
		result := envTestResult{Level: int(GetLevel())}
		data, _ := json.Marshal(result)
		os.Stdout.Write(data)
		os.Exit(0)
		return
	}

	tests := []struct {
		name     string
		appEnv   string
		wantLevel Level
	}{
		{"dev", "dev", DebugLevel},
		{"development", "development", DebugLevel},
		{"DEV_uppercase", "DEV", DebugLevel},
		{"Development_mixed", "Development", DebugLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runEnvSubprocess(t, "TestEnvInit_DevEnvironment", tt.appEnv)
			if got != tt.wantLevel {
				t.Errorf("APP_ENV=%s: expected level %v, got %v", tt.appEnv, tt.wantLevel, got)
			}
		})
	}
}

func TestEnvInit_TestEnvironment(t *testing.T) {
	if os.Getenv("TEST_ENV_INIT") == "1" {
		result := envTestResult{Level: int(GetLevel())}
		data, _ := json.Marshal(result)
		os.Stdout.Write(data)
		os.Exit(0)
		return
	}

	tests := []struct {
		name     string
		appEnv   string
		wantLevel Level
	}{
		{"test", "test", DebugLevel},
		{"canary", "canary", DebugLevel},
		{"TEST_uppercase", "TEST", DebugLevel},
		{"Canary_mixed", "Canary", DebugLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runEnvSubprocess(t, "TestEnvInit_TestEnvironment", tt.appEnv)
			if got != tt.wantLevel {
				t.Errorf("APP_ENV=%s: expected level %v, got %v", tt.appEnv, tt.wantLevel, got)
			}
		})
	}
}

func TestEnvInit_ProdEnvironment(t *testing.T) {
	if os.Getenv("TEST_ENV_INIT") == "1" {
		result := envTestResult{Level: int(GetLevel())}
		data, _ := json.Marshal(result)
		os.Stdout.Write(data)
		os.Exit(0)
		return
	}

	tests := []struct {
		name     string
		appEnv   string
		wantLevel Level
	}{
		{"prod", "prod", InfoLevel},
		{"release", "release", InfoLevel},
		{"production", "production", InfoLevel},
		{"PROD_uppercase", "PROD", InfoLevel},
		{"Release_mixed", "Release", InfoLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runEnvSubprocess(t, "TestEnvInit_ProdEnvironment", tt.appEnv)
			if got != tt.wantLevel {
				t.Errorf("APP_ENV=%s: expected level %v, got %v", tt.appEnv, tt.wantLevel, got)
			}
		})
	}
}

func TestEnvInit_DefaultEnvironment(t *testing.T) {
	if os.Getenv("TEST_ENV_INIT") == "1" {
		result := envTestResult{Level: int(GetLevel())}
		data, _ := json.Marshal(result)
		os.Stdout.Write(data)
		os.Exit(0)
		return
	}

	tests := []struct {
		name   string
		appEnv string
	}{
		{"empty", ""},
		{"unknown", "unknown"},
		{"staging", "staging"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runEnvSubprocess(t, "TestEnvInit_DefaultEnvironment", tt.appEnv)
			// Default level is DebugLevel (set in newLogger)
			if got != DebugLevel {
				t.Errorf("APP_ENV=%q: expected default level %v, got %v", tt.appEnv, DebugLevel, got)
			}
		})
	}
}
