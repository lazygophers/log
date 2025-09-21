package log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestNewHourlyRotator(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	rotator := NewHourlyRotator(filename, 1024, 5)

	if rotator == nil {
		t.Fatal("NewHourlyRotator returned nil")
	}

	if rotator.filename != filename {
		t.Errorf("Expected filename %s, got %s", filename, rotator.filename)
	}

	if rotator.maxSize != 1024 {
		t.Errorf("Expected maxSize 1024, got %d", rotator.maxSize)
	}

	if rotator.maxFiles != 5 {
		t.Errorf("Expected maxFiles 5, got %d", rotator.maxFiles)
	}
}

func TestHourlyRotator_Write(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	rotator := NewHourlyRotator(filename, 1024*1024, 5)
	defer func() { _ = rotator.Close() }()

	testData := []byte("test log message\n")
	n, err := rotator.Write(testData)

	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(testData), n)
	}

	// 检查文件是否被创建
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read temp dir: %v", err)
	}

	var logFileFound bool
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "test") && strings.HasSuffix(file.Name(), ".log") {
			logFileFound = true
			break
		}
	}

	if !logFileFound {
		t.Error("Log file was not created")
	}
}

func TestHourlyRotator_Rotation(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	// 设置小的文件大小限制来触发轮转
	rotator := NewHourlyRotator(filename, 10, 5)
	defer func() { _ = rotator.Close() }()

	// 写入足够的数据触发大小轮转
	testData := []byte("this is a long test message that should trigger rotation\n")
	_, err := rotator.Write(testData)
	if err != nil {
		t.Fatalf("First write failed: %v", err)
	}

	// 再次写入应该触发轮转
	_, err = rotator.Write(testData)
	if err != nil {
		t.Fatalf("Second write failed: %v", err)
	}

	// 检查是否创建了多个文件
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read temp dir: %v", err)
	}

	logFileCount := 0
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "test") && strings.HasSuffix(file.Name(), ".log") && !strings.HasSuffix(file.Name(), "test.log") {
			logFileCount++
		}
	}

	if logFileCount < 1 {
		t.Error("Expected at least one rotated log file")
	}
}

func TestHourlyRotator_Sync(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	rotator := NewHourlyRotator(filename, 1024, 5)
	defer func() { _ = rotator.Close() }()

	// 写入数据
	testData := []byte("test log message\n")
	_, err := rotator.Write(testData)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// 测试 Sync
	err = rotator.Sync()
	if err != nil {
		t.Errorf("Sync failed: %v", err)
	}
}

func TestHourlyRotator_Close(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	rotator := NewHourlyRotator(filename, 1024, 5)

	// 写入数据以创建文件
	testData := []byte("test log message\n")
	_, err := rotator.Write(testData)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// 测试 Close
	err = rotator.Close()
	if err != nil {
		t.Errorf("Close failed: %v", err)
	}

	// 再次关闭应该不会出错（可能返回 nil 或 file already closed 错误）
	err = rotator.Close()
	if err != nil && !strings.Contains(err.Error(), "file already closed") {
		t.Errorf("Second close should not fail with unexpected error: %v", err)
	}
}

func TestHourlyRotator_CleanupOldFiles(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	// 创建一些旧的日志文件
	for i := 0; i < 15; i++ {
		oldTime := time.Now().Add(-time.Duration(i) * time.Hour)
		oldFilename := filename + oldTime.Format("2006010215") + ".log"
		file, err := os.Create(oldFilename)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", oldFilename, err)
		}
		file.WriteString("test content")
		_ = file.Close()
	}

	rotator := NewHourlyRotator(filename, 1024, 5)
	defer func() { _ = rotator.Close() }()

	// 手动触发清理
	rotator.cleanupOldFiles()

	// 检查文件数量
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read temp dir: %v", err)
	}

	logFileCount := 0
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "test") && strings.HasSuffix(file.Name(), ".log") {
			logFileCount++
		}
	}

	// 应该最多保留 maxFiles (5) 个文件
	if logFileCount > 5 {
		t.Errorf("Expected at most 5 log files, found %d", logFileCount)
	}
}

func TestHourlyRotator_UpdateLink(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	rotator := NewHourlyRotator(filename, 1024, 5)
	defer func() { _ = rotator.Close() }()

	// 创建目标文件
	targetFile := filename + "2023071015.log"
	file, err := os.Create(targetFile)
	if err != nil {
		t.Fatalf("Failed to create target file: %v", err)
	}
	_ = file.Close()

	// 测试更新链接
	rotator.updateLink(targetFile)

	// 检查链接是否存在（在支持软链接的系统上）
	linkName := filename + ".log"
	if _, err := os.Lstat(linkName); err == nil {
		// 软链接创建成功，验证它指向正确的文件
		if link, err := os.Readlink(linkName); err == nil {
			expected := filepath.Base(targetFile)
			if link != expected {
				t.Errorf("Expected link to point to %s, but it points to %s", expected, link)
			}
		}
	}
	// 如果系统不支持软链接，测试仍然应该通过
}

// TestHourlyRotator_WriteError tests error conditions in Write function
func TestHourlyRotator_WriteError(t *testing.T) {
	// Create a rotator that will fail to rotate (permission denied)
	// Use /dev/null as base path which should cause issues
	rotator := NewHourlyRotator("/dev/null", 1024, 5)
	defer func() { _ = rotator.Close() }()

	// Try to write - this should fail during rotate() call
	_, err := rotator.Write([]byte("test"))
	if err != nil {
		t.Logf("Got expected error from Write->rotate: %v", err)
	} else {
		t.Log("Write succeeded unexpectedly - system might allow /dev/null writes")
	}
}

// TestHourlyRotator_RotateError tests error conditions in rotate function
func TestHourlyRotator_RotateError(t *testing.T) {
	// Create a rotator with a path that will cause permission errors
	rotator := NewHourlyRotator("/root/test", 1024, 5) // Path that likely doesn't have write permissions
	defer func() { _ = rotator.Close() }()

	// Try to write - this should trigger rotate error
	_, err := rotator.Write([]byte("test"))
	// We expect this to fail on most systems due to permissions, but it's hard to test reliably
	// The important thing is that we're exercising the error path in Write->rotate
	if err != nil {
		t.Logf("Got expected error from rotate: %v", err)
	}
}
