package log

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSetOutput(t *testing.T) {
	// 保存原始的标准日志记录器状态
	originalOut := std.out
	defer func() {
		std.out = originalOut
	}()

	var buf1, buf2 bytes.Buffer
	result := SetOutput(&buf1, &buf2)

	if result != std {
		t.Error("SetOutput should return the global std logger")
	}

	// 测试写入多个输出
	Info("test message")

	if buf1.Len() == 0 {
		t.Error("First buffer should contain data")
	}

	if buf2.Len() == 0 {
		t.Error("Second buffer should contain data")
	}

	// 两个 buffer 的内容应该相同
	if buf1.String() != buf2.String() {
		t.Error("Both buffers should contain the same data")
	}
}

func TestGetOutputWriter(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.log")

	writer := GetOutputWriter(filename)

	if writer == nil {
		t.Fatal("GetOutputWriter returned nil")
	}

	// 测试写入
	testData := []byte("test log message\n")
	n, err := writer.Write(testData)

	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(testData), n)
	}

	// 检查文件是否存在并包含正确内容
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Failed to read log file: %v", err)
	}

	if string(content) != string(testData) {
		t.Errorf("Expected file content %q, got %q", string(testData), string(content))
	}
}

func TestGetOutputWriter_DirectoryCreation(t *testing.T) {
	tmpDir := t.TempDir()
	// 创建一个不存在的子目录路径
	filename := filepath.Join(tmpDir, "subdir", "nested", "test.log")

	writer := GetOutputWriter(filename)

	if writer == nil {
		t.Fatal("GetOutputWriter returned nil")
	}

	// 测试写入（这应该创建目录）
	testData := []byte("test log message\n")
	_, err := writer.Write(testData)

	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	// 检查目录是否被创建
	if _, err := os.Stat(filepath.Dir(filename)); os.IsNotExist(err) {
		t.Error("Directory should have been created")
	}

	// 检查文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Log file should have been created")
	}
}

func TestGetOutputWriterHourly(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	writer := GetOutputWriterHourly(filename)

	if writer == nil {
		t.Fatal("GetOutputWriterHourly returned nil")
	}

	// 测试写入
	testData := []byte("test log message\n")
	n, err := writer.Write(testData)

	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(testData), n)
	}

	// 检查是否创建了按小时命名的日志文件
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read temp dir: %v", err)
	}

	var logFileFound bool
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, "test") && strings.HasSuffix(name, ".log") && len(name) > 10 {
			logFileFound = true
			break
		}
	}

	if !logFileFound {
		t.Error("Hourly log file was not created")
	}
}

func TestGetOutputWriterHourly_Reuse(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test")

	// 获取两次相同文件名的 writer
	writer1 := GetOutputWriterHourly(filename)
	writer2 := GetOutputWriterHourly(filename)

	// 应该返回相同的实例（缓存）
	if writer1 != writer2 {
		t.Error("GetOutputWriterHourly should reuse instances for same filename")
	}
}

func TestEnsureDir(t *testing.T) {
	tmpDir := t.TempDir()
	testDir := filepath.Join(tmpDir, "test", "nested", "dir")

	// 目录不存在
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Fatal("Test directory should not exist initially")
	}

	// 调用 ensureDir
	ensureDir(testDir)

	// 目录应该被创建
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Error("Directory should have been created")
	}

	// 检查目录是否真的是目录
	info, err := os.Stat(testDir)
	if err != nil {
		t.Errorf("Failed to stat directory: %v", err)
	}

	if !info.IsDir() {
		t.Error("Created path should be a directory")
	}
}

func TestEnsureDir_CurrentDirectory(t *testing.T) {
	// 测试当目录是 "." 时的行为
	ensureDir(".")
	// 这不应该做任何事情，也不应该出错
}

func TestEnsureDir_ExistingDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	testDir := filepath.Join(tmpDir, "existing")

	// 先创建目录
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// 调用 ensureDir 应该不会出错
	ensureDir(testDir)

	// 目录应该仍然存在
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Error("Existing directory should still exist")
	}
}

func TestIsDir(t *testing.T) {
	tmpDir := t.TempDir()

	// 测试目录
	testDir := filepath.Join(tmpDir, "testdir")
	err := os.Mkdir(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	if !isDir(testDir) {
		t.Error("isDir should return true for directory")
	}

	// 测试文件
	testFile := filepath.Join(tmpDir, "testfile.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()

	if isDir(testFile) {
		t.Error("isDir should return false for file")
	}

	// 测试不存在的路径
	nonExistent := filepath.Join(tmpDir, "nonexistent")
	if isDir(nonExistent) {
		t.Error("isDir should return false for non-existent path")
	}
}

func TestOutput_Interface(t *testing.T) {
	// 测试 Output 接口
	var buf bytes.Buffer

	// bytes.Buffer 实现了 io.Writer，所以可以用作 Output
	var output Output = &buf

	testData := []byte("test data")
	n, err := output.Write(testData)

	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(testData), n)
	}

	if buf.String() != string(testData) {
		t.Errorf("Expected %q, got %q", string(testData), buf.String())
	}
}

// 测试并发安全性
func TestRotatorInstances_ConcurrentAccess(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "concurrent")

	// 清理全局状态
	rotatorMutex.Lock()
	delete(rotatorInstances, filename)
	initialCount := len(rotatorInstances)
	rotatorMutex.Unlock()

	// 启动多个 goroutine 同时访问
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
			defer func() { done <- true }()
			writer := GetOutputWriterHourly(filename)
			if writer == nil {
				t.Error("GetOutputWriterHourly returned nil")
			}
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}

	// 应该只有一个新实例被创建
	rotatorMutex.Lock()
	finalCount := len(rotatorInstances)
	rotatorMutex.Unlock()

	expectedCount := initialCount + 1
	if finalCount != expectedCount {
		t.Errorf("Expected %d rotator instances, got %d", expectedCount, finalCount)
	}
}

// TestEnsureDir_CreateError tests the error case when directory creation fails
func TestEnsureDir_CreateError(t *testing.T) {
	// Create a path that will fail to create (e.g., under a non-writable directory)
	// This is difficult to test portably, so we'll test the "." case which is already covered
	// and the existing directory case which is also covered.
	// The error branch is difficult to test without platform-specific code.

	// Test that ensureDir handles current directory gracefully
	ensureDir(".")
	// Should not panic or error

	// Test empty string (edge case)
	ensureDir("")
	// Should not panic
}
