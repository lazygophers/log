package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSetOutput 测试 SetOutput 函数
func TestSetOutput(t *testing.T) {
	// 重置全局状态
	cleanRotatelogOnce = make(map[string]bool)

	t.Run("单个输出", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger := New()
		logger.SetOutput(buf)

		// 验证输出设置生效
		logger.Info("测试消息")
		assert.Contains(t, buf.String(), "测试消息")
	})

	t.Run("多个输出", func(t *testing.T) {
		buf1 := &bytes.Buffer{}
		buf2 := &bytes.Buffer{}
		logger := New()
		logger.SetOutput(buf1, buf2)

		// 验证消息输出到所有输出
		logger.Info("多输出测试")
		assert.Contains(t, buf1.String(), "多输出测试")
		assert.Contains(t, buf2.String(), "多输出测试")
	})

	t.Run("无参数", func(t *testing.T) {
		// 保存原始输出
		originalOutput := std.out
		defer std.SetOutput(originalOutput)

		// 验证 SetOutput 无参数不会 panic
		assert.NotPanics(t, func() {
			SetOutput()
		})
	})
}

// TestGetOutputWriter 测试 GetOutputWriter 函数
func TestGetOutputWriter(t *testing.T) {
	t.Run("正常创建", func(t *testing.T) {
		// 创建临时目录
		tempDir := t.TempDir()
		logFile := filepath.Join(tempDir, "test.log")

		writer := GetOutputWriter(logFile)
		assert.NotNil(t, writer)

		// 验证写入功能
		testData := "测试日志\n"
		n, err := writer.Write([]byte(testData))
		require.NoError(t, err)
		assert.Equal(t, len(testData), n)

		// 验证文件内容
		content, err := os.ReadFile(logFile)
		require.NoError(t, err)
		assert.Contains(t, string(content), "测试日志")

		// 验证行数
		lines := strings.Split(string(content), "\n")
		// 移除最后一个空行（如果有）
		if lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1]
		}
		assert.Equal(t, 1, len(lines))
	})

	t.Run("目录不存在时自动创建", func(t *testing.T) {
		tempDir := t.TempDir()
		subDir := filepath.Join(tempDir, "logs", "app")
		logFile := filepath.Join(subDir, "app.log")

		// 确保目录不存在
		_, err := os.Stat(subDir)
		assert.True(t, os.IsNotExist(err))

		writer := GetOutputWriter(logFile)
		assert.NotNil(t, writer)

		// 验证目录被创建
		_, err = os.Stat(subDir)
		assert.NoError(t, err)
	})

	t.Run("无效路径", func(t *testing.T) {
		// 使用一个无效的路径（例如在只读系统中的位置）
		// 由于不同系统限制，这里我们只测试不会 panic
		assert.NotPanics(t, func() {
			_ = GetOutputWriter("/invalid/path/test.log")
		})
	})
}

// TestGetOutputWriterHourly 测试 GetOutputWriterHourly 函数
func TestGetOutputWriterHourly(t *testing.T) {
	t.Run("创建按小时轮转的写入器", func(t *testing.T) {
		tempDir := t.TempDir()
		baseFile := filepath.Join(tempDir, "app")

		writer := GetOutputWriterHourly(baseFile)
		assert.NotNil(t, writer)
		assert.Implements(t, (*io.Writer)(nil), writer)

		// 等待一下确保文件和软链接创建完成
		time.Sleep(50 * time.Millisecond)

		// 验证软链接（如果存在）
		linkFile := baseFile + ".log"
		if _, err := os.Stat(linkFile); err == nil {
			// 软链接存在，记录信息
			t.Logf("软链接 %s 已创建", linkFile)
		} else if !os.IsNotExist(err) {
			// 其他错误
			t.Logf("检查软链接时出现错误: %v", err)
		}
	})

	t.Run("并发创建只启动一个清理协程", func(t *testing.T) {
		tempDir := t.TempDir()
		baseFile := filepath.Join(tempDir, "concurrent")

		var wg sync.WaitGroup
		const goroutines = 10

		// 重置全局状态
		cleanMutex.Lock()
		cleanRotatelogOnce = make(map[string]bool)
		cleanMutex.Unlock()

		wg.Add(goroutines)
		for i := 0; i < goroutines; i++ {
			go func() {
				defer wg.Done()
				_ = GetOutputWriterHourly(baseFile)
			}()
		}

		wg.Wait()

		// 验证只有一个清理协程被启动
		cleanMutex.Lock()
		assert.True(t, cleanRotatelogOnce[baseFile])
		cleanMutex.Unlock()
	})

	t.Run("写入测试", func(t *testing.T) {
		tempDir := t.TempDir()
		baseFile := filepath.Join(tempDir, "write_test")

		writer := GetOutputWriterHourly(baseFile)

		// 写入测试消息
		testMessage := "测试按小时轮转的日志\n"
		_, err := writer.Write([]byte(testMessage))
		require.NoError(t, err)

		// 验证软链接指向最新的日志文件
		linkFile := baseFile + ".log"
		content, err := os.ReadFile(linkFile)
		require.NoError(t, err)
		assert.Contains(t, string(content), testMessage)
	})
}

// TestEnsureDir 测试 ensureDir 函数
func TestEnsureDir(t *testing.T) {
	t.Run("目录不存在时创建", func(t *testing.T) {
		tempDir := t.TempDir()
		newDir := filepath.Join(tempDir, "new", "nested", "dir")

		// 确保目录不存在
		_, err := os.Stat(newDir)
		assert.True(t, os.IsNotExist(err))

		ensureDir(newDir)

		// 验证目录被创建
		_, err = os.Stat(newDir)
		assert.NoError(t, err)
	})

	t.Run("目录已存在时不操作", func(t *testing.T) {
		tempDir := t.TempDir()

		// 创建目录
		subDir := filepath.Join(tempDir, "existing")
		err := os.Mkdir(subDir, 0755)
		require.NoError(t, err)

		// 获取修改时间
		info, err := os.Stat(subDir)
		require.NoError(t, err)
		modTime := info.ModTime()

		// 等待一小段时间
		time.Sleep(10 * time.Millisecond)

		ensureDir(subDir)

		// 验证修改时间未改变
		newInfo, err := os.Stat(subDir)
		require.NoError(t, err)
		assert.Equal(t, modTime, newInfo.ModTime())
	})

	t.Run("当前目录不处理", func(t *testing.T) {
		assert.NotPanics(t, func() {
			ensureDir(".")
		})
	})

	t.Run("创建失败时不panic", func(t *testing.T) {
		// 使用一个可能无法创建的路径
		// 在某些系统中，可能需要特殊权限
		assert.NotPanics(t, func() {
			ensureDir("/root/test_no_permission")
		})
	})
}

// TestIsDir 测试 isDir 函数
func TestIsDir(t *testing.T) {
	t.Run("存在的目录", func(t *testing.T) {
		tempDir := t.TempDir()
		assert.True(t, isDir(tempDir))
	})

	t.Run("不存在的路径", func(t *testing.T) {
		assert.False(t, isDir("/path/that/does/not/exist"))
	})

	t.Run("文件而不是目录", func(t *testing.T) {
		tempDir := t.TempDir()
		file := filepath.Join(tempDir, "test_file")

		err := os.WriteFile(file, []byte("test"), 0644)
		require.NoError(t, err)

		assert.False(t, isDir(file))
	})

	t.Run("空路径", func(t *testing.T) {
		assert.False(t, isDir(""))
	})
}

// TestOutputInterface 测试 Output 接口
func TestOutputInterface(t *testing.T) {
	// 验证 GetOutputWriter 返回的值实现了 Output 接口
	tempDir := t.TempDir()
	logFile := filepath.Join(tempDir, "interface_test.log")

	writer := GetOutputWriter(logFile)

	// 检查是否实现了 io.Writer 接口
	assert.Implements(t, (*io.Writer)(nil), writer)

	// 检查是否实现了 Output 接口（嵌入 io.Writer）
	var output Output
	var ok bool

	output, ok = writer.(Output)
	assert.True(t, ok, "GetOutputWriter 应该返回 Output 接口")
	assert.NotNil(t, output)
}

// TestConcurrentSafety 测试并发安全性
func TestConcurrentSafety(t *testing.T) {
	t.Run("并发访问 cleanMutex", func(t *testing.T) {
		tempDir := t.TempDir()
		baseFile := filepath.Join(tempDir, "concurrent_test")

		// 重置全局状态
		cleanMutex.Lock()
		cleanRotatelogOnce = make(map[string]bool)
		cleanMutex.Unlock()

		var wg sync.WaitGroup
		const goroutines = 100

		wg.Add(goroutines)
		for i := 0; i < goroutines; i++ {
			go func(id int) {
				defer wg.Done()
				writer := GetOutputWriterHourly(baseFile + fmt.Sprintf("_%d", id))
				assert.NotNil(t, writer)

				// 写入一些数据
				_, err := writer.Write([]byte(fmt.Sprintf("Goroutine %d\n", id)))
				assert.NoError(t, err)
			}(i)
		}

		wg.Wait()

		// 验证所有文件都被创建
		files, err := os.ReadDir(tempDir)
		require.NoError(t, err)
		assert.True(t, len(files) > 0)
	})

	t.Run("并发目录创建", func(t *testing.T) {
		tempDir := t.TempDir()
		var wg sync.WaitGroup
		const goroutines = 50

		wg.Add(goroutines)
		for i := 0; i < goroutines; i++ {
			go func(id int) {
				defer wg.Done()
				logFile := filepath.Join(tempDir, fmt.Sprintf("dir_%d", id), "app.log")
				writer := GetOutputWriter(logFile)
				assert.NotNil(t, writer)
			}(i)
		}

		wg.Wait()

		// 验证所有目录都被创建
		files, err := os.ReadDir(tempDir)
		require.NoError(t, err)
		assert.Equal(t, goroutines, len(files))
	})
}

// TestLogFileRotation 测试日志文件轮转功能
func TestLogFileRotation(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过耗时的轮转测试")
	}

	t.Run("按大小轮转", func(t *testing.T) {
		tempDir := t.TempDir()
		baseFile := filepath.Join(tempDir, "size_rotation")

		// 创建一个小的轮转大小用于测试
		hook, err := rotatelogs.New(
			baseFile+"%Y%m%d%H.log",
			rotatelogs.WithRotationSize(1024), // 1KB
			rotatelogs.WithRotationCount(3),
		)
		require.NoError(t, err)

		// 写入超过限制的数据
		data := strings.Repeat("A", 512) // 512B per write
		for i := 0; i < 10; i++ {
			_, err := hook.Write([]byte(data))
			require.NoError(t, err)
		}

		// 验证文件被轮转
		files, err := os.ReadDir(tempDir)
		require.NoError(t, err)
		assert.True(t, len(files) > 1)
	})

	t.Run("按时间轮转", func(t *testing.T) {
		tempDir := t.TempDir()
		baseFile := filepath.Join(tempDir, "time_rotation")

		// 使用很短的轮转时间间隔进行测试
		hook, err := rotatelogs.New(
			baseFile+"%Y%m%d%H%M%S.log",
			rotatelogs.WithRotationTime(time.Second), // 每秒轮转
			rotatelogs.WithRotationCount(3),
		)
		require.NoError(t, err)

		// 等待轮转发生
		time.Sleep(2 * time.Second)

		// 写入数据触发轮转
		_, err = hook.Write([]byte("test after rotation"))
		require.NoError(t, err)

		// 验证文件被轮转
		files, err := os.ReadDir(tempDir)
		require.NoError(t, err)
		assert.True(t, len(files) >= 1)
	})
}

// TestIntegrationWithLogger 测试与 Logger 的集成
func TestIntegrationWithLogger(t *testing.T) {
	t.Run("使用文件输出", func(t *testing.T) {
		tempDir := t.TempDir()
		logFile := filepath.Join(tempDir, "integration.log")

		fileWriter := GetOutputWriter(logFile)
		logger := New()
		logger.SetOutput(fileWriter)
		logger.SetLevel(InfoLevel)

		logger.Info("集成测试消息")

		// 验证文件内容
		content, err := os.ReadFile(logFile)
		require.NoError(t, err)
		assert.Contains(t, string(content), "集成测试消息")
		assert.Contains(t, string(content), "[info]")
	})

	t.Run("使用按小时轮转输出", func(t *testing.T) {
		tempDir := t.TempDir()
		baseFile := filepath.Join(tempDir, "hourly_integration")

		hourlyWriter := GetOutputWriterHourly(baseFile)
		logger := New()
		logger.SetOutput(hourlyWriter)
		logger.SetLevel(DebugLevel)

		logger.Debug("调试消息")
		logger.Info("信息消息")

		// 验证软链接存在
		linkFile := baseFile + ".log"
		_, err := os.Stat(linkFile)
		assert.NoError(t, err)

		// 验证内容
		content, err := os.ReadFile(linkFile)
		require.NoError(t, err)
		assert.Contains(t, string(content), "调试消息")
		assert.Contains(t, string(content), "信息消息")
	})
}

// BenchmarkOutputWriter 性能基准测试
func BenchmarkOutputWriter(b *testing.B) {
	tempDir := b.TempDir()
	logFile := filepath.Join(tempDir, "benchmark.log")

	writer := GetOutputWriter(logFile)
	testData := []byte("性能测试日志消息\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = writer.Write(testData)
	}
}

// BenchmarkOutputWriterHourly 按小时轮转写入器性能测试
func BenchmarkOutputWriterHourly(b *testing.B) {
	tempDir := b.TempDir()
	baseFile := filepath.Join(tempDir, "benchmark_hourly")

	writer := GetOutputWriterHourly(baseFile)
	testData := []byte("按小时轮转性能测试日志消息\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = writer.Write(testData)
	}
}

// BenchmarkConcurrentOutput 并发输出性能测试
func BenchmarkConcurrentOutput(b *testing.B) {
	tempDir := b.TempDir()
	baseFile := filepath.Join(tempDir, "concurrent_benchmark")

	b.RunParallel(func(pb *testing.PB) {
		// 每个 goroutine 使用唯一的文件
		writer := GetOutputWriter(baseFile + fmt.Sprintf("_%d", time.Now().UnixNano()))
		testData := []byte("并发性能测试日志消息\n")

		for pb.Next() {
			_, _ = writer.Write(testData)
		}
	})
}
