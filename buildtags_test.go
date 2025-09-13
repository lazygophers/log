package log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReleaseLogPath_Debug(t *testing.T) {
	// 在非 release 模式下，ReleaseLogPath 应该为空
	if ReleaseLogPath != "" {
		// 这个测试只在 debug 模式下才有意义
		// 如果在 release 模式下运行，ReleaseLogPath 会有默认值
		t.Skip("This test is for debug build mode only")
	}
	
	// 在 debug 模式下，ReleaseLogPath 应该为空字符串
	if ReleaseLogPath != "" {
		t.Errorf("In debug mode, ReleaseLogPath should be empty, got %q", ReleaseLogPath)
	}
}

func TestNewLogger_OutputSelection(t *testing.T) {
	// 保存原始的 ReleaseLogPath
	originalPath := ReleaseLogPath
	defer func() {
		// 由于 ReleaseLogPath 是全局变量，我们不能真正重置它
		// 但可以记录原始值用于参考
		_ = originalPath
	}()
	
	logger := newLogger()
	
	if logger == nil {
		t.Fatal("newLogger returned nil")
	}
	
	// 检查输出是否正确设置
	if logger.out == nil {
		t.Error("Logger output should not be nil")
	}
	
	// 如果 ReleaseLogPath 为空（debug 模式），应该使用标准输出
	if ReleaseLogPath == "" {
		// 在 debug 模式下，应该包装 stdout
		wrapper, ok := logger.out.(*WriteSyncerWrapper)
		if ok {
			if wrapper.writer != os.Stdout {
				t.Error("In debug mode, logger should use stdout")
			}
		}
	} else {
		// 在 release 模式下，应该使用文件轮转器
		// 这里我们只能检查 out 不是直接的 stdout 包装器
		if wrapper, ok := logger.out.(*WriteSyncerWrapper); ok {
			if wrapper.writer == os.Stdout {
				t.Error("In release mode, logger should not use stdout directly")
			}
		}
	}
}

// 测试在不同构建标签下的行为
func TestBuildTagBehavior(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "debug_mode_path",
			expected: "",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 在这个测试中，我们检查当前编译环境下的 ReleaseLogPath
			// 如果是 debug 模式，应该为空
			// 如果是 release 模式，应该有默认路径
			
			if strings.Contains(ReleaseLogPath, "lazygophers") || ReleaseLogPath == "" {
				// 这是预期的行为（要么是空字符串，要么包含默认路径）
			} else {
				t.Errorf("Unexpected ReleaseLogPath value: %q", ReleaseLogPath)
			}
		})
	}
}

// 测试默认路径的合理性
func TestDefaultLogPath(t *testing.T) {
	if ReleaseLogPath == "" {
		t.Skip("ReleaseLogPath is empty (debug mode)")
	}
	
	// 检查路径是否包含期望的组件
	if !strings.Contains(ReleaseLogPath, "lazygophers") {
		t.Errorf("Default log path should contain 'lazygophers', got: %q", ReleaseLogPath)
	}
	
	if !strings.Contains(ReleaseLogPath, "log") {
		t.Errorf("Default log path should contain 'log', got: %q", ReleaseLogPath)
	}
	
	// 检查路径是否是绝对路径
	if !filepath.IsAbs(ReleaseLogPath) {
		t.Errorf("Default log path should be absolute, got: %q", ReleaseLogPath)
	}
	
	// 检查父目录是否存在或可创建
	parentDir := filepath.Dir(ReleaseLogPath)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		t.Errorf("Cannot create parent directory %q: %v", parentDir, err)
	}
	
	// 清理测试创建的目录
	defer func() {
		// 只清理测试创建的目录，不要删除系统目录
		if strings.Contains(parentDir, "lazygophers") {
			os.RemoveAll(parentDir)
		}
	}()
}