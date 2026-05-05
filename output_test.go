package log

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func TestGetOutputWriterEdgeCases(t *testing.T) {
	t.Run("GetOutputWriter_creates_parent_dirs", func(t *testing.T) {
		tmpDir := t.TempDir()
		deepPath := filepath.Join(tmpDir, "level1", "level2", "level3", "test.log")

		writer := GetOutputWriter(deepPath)
		if writer == nil {
			t.Fatal("GetOutputWriter should return writer")
		}

		// Verify parent directory was created
		parentDir := filepath.Join(tmpDir, "level1", "level2", "level3")
		if _, err := os.Stat(parentDir); os.IsNotExist(err) {
			t.Errorf("Parent directory %s should be created", parentDir)
		}
	})

	t.Run("GetOutputWriter_same_file_twice", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := filepath.Join(tmpDir, "test.log")

		writer1 := GetOutputWriter(logFile)
		writer2 := GetOutputWriter(logFile)

		if writer1 == nil || writer2 == nil {
			t.Error("GetOutputWriter should return writer")
		}

		// Both writers should work
		writer1.Write([]byte("line1\n"))
		writer2.Write([]byte("line2\n"))

		// Verify file exists and has content
		content, err := os.ReadFile(logFile)
		if err != nil {
			t.Fatalf("Failed to read log file: %v", err)
		}

		contentStr := string(content)
		if !strings.Contains(contentStr, "line1") || !strings.Contains(contentStr, "line2") {
			t.Error("Both writers should write to file")
		}
	})
}

func TestGetOutputWriterHourly(t *testing.T) {
	t.Run("GetOutputWriterHourly_returns_cached_instance", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Clear any existing instances for this test
		rotatorInstances = make(map[string]*HourlyRotator)
		rotatorMutex = &sync.Mutex{}

		writer1 := GetOutputWriterHourly(tmpDir)
		writer2 := GetOutputWriterHourly(tmpDir)

		if writer1 != writer2 {
			t.Error("GetOutputWriterHourly should return cached instance")
		}

		// Cleanup
		if rotator, ok := writer1.(*HourlyRotator); ok {
			rotator.Close()
		}
	})

	t.Run("GetOutputWriterHourly_creates_new_instance", func(t *testing.T) {
		tmpDir := t.TempDir()
		otherDir := t.TempDir()

		// Clear any existing instances for this test
		rotatorInstances = make(map[string]*HourlyRotator)
		rotatorMutex = &sync.Mutex{}

		writer1 := GetOutputWriterHourly(tmpDir)
		writer2 := GetOutputWriterHourly(otherDir)

		if writer1 == writer2 {
			t.Error("GetOutputWriterHourly should return different instances for different paths")
		}

		// Cleanup
		if rotator, ok := writer1.(*HourlyRotator); ok {
			rotator.Close()
		}
		if rotator, ok := writer2.(*HourlyRotator); ok {
			rotator.Close()
		}
	})
}

func TestEnsureDirCoverage(t *testing.T) {
	t.Run("ensureDir_with_dot_path", func(t *testing.T) {
		// Test the "." path branch (should not create directory)
		// This is implicitly tested when creating files in current directory
		tmpDir := t.TempDir()
		oldDir, _ := os.Getwd()
		defer os.Chdir(oldDir)

		os.Chdir(tmpDir)
		writer := GetOutputWriter("test.log")
		if writer == nil {
			t.Error("GetOutputWriter should work with current directory")
		}
	})

	t.Run("ensureDir_with_existing_directory", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Directory already exists from TempDir
		writer := GetOutputWriter(tmpDir + "/test.log")
		if writer == nil {
			t.Error("GetOutputWriter should work with existing directory")
		}
	})

	t.Run("GetOutputWriter_nested_paths", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Test deeply nested path creation
		deepPath := tmpDir + "/a/b/c/d/test.log"
		writer := GetOutputWriter(deepPath)
		if writer == nil {
			t.Error("GetOutputWriter should create nested directories")
		}

		// Verify all directories were created
		dirs := []string{
			tmpDir + "/a",
			tmpDir + "/a/b",
			tmpDir + "/a/b/c",
			tmpDir + "/a/b/c/d",
		}
		for _, dir := range dirs {
			if info, err := os.Stat(dir); err != nil || !info.IsDir() {
				t.Errorf("Directory %s should exist", dir)
			}
		}
	})
}
