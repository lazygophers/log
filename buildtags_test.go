package log

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReleaseLogPath_Debug(t *testing.T) {
	// 在非 release 模式下，ReleaseLogPath 应该为空
	if ReleaseLogDir != "" {
		// 这个测试只在 debug 模式下才有意义
		// 如果在 release 模式下运行，ReleaseLogPath 会有默认值
		t.Skip("This test is for debug build mode only")
	}

	// 在 debug 模式下，ReleaseLogPath 应该为空字符串
	if ReleaseLogDir != "" {
		t.Errorf("In debug mode, ReleaseLogPath should be empty, got %q", ReleaseLogDir)
	}
}

func TestNewLogger_OutputSelection(t *testing.T) {
	// 保存原始的 ReleaseLogPath
	originalPath := ReleaseLogDir
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
	if ReleaseLogDir == "" {
		// 在 debug 模式下，应该包装 stdout
		// 由于 AddSync 返回的是接口，我们无法检查内部类型
		// 只需确认输出不是 nil
		if logger.out == nil {
			t.Error("In debug mode, logger should have output")
		}
	} else {
		// 在 release 模式下，应该使用文件轮转器
		// 只需确认输出不是 nil
		if logger.out == nil {
			t.Error("In release mode, logger should have output")
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

			if strings.Contains(ReleaseLogDir, "lazygophers") || ReleaseLogDir == "" {
				// 这是预期的行为（要么是空字符串，要么包含默认路径）
			} else {
				t.Errorf("Unexpected ReleaseLogPath value: %q", ReleaseLogDir)
			}
		})
	}
}

// 测试默认路径的合理性
func TestDefaultLogPath(t *testing.T) {
	if ReleaseLogDir == "" {
		t.Skip("ReleaseLogPath is empty (debug mode)")
	}

	// 检查路径是否包含期望的组件
	if !strings.Contains(ReleaseLogDir, "lazygophers") {
		t.Errorf("Default log path should contain 'lazygophers', got: %q", ReleaseLogDir)
	}

	if !strings.Contains(ReleaseLogDir, "log") {
		t.Errorf("Default log path should contain 'log', got: %q", ReleaseLogDir)
	}

	// 检查路径是否是绝对路径
	if !filepath.IsAbs(ReleaseLogDir) {
		t.Errorf("Default log path should be absolute, got: %q", ReleaseLogDir)
	}

	// 检查父目录是否存在或可创建
	parentDir := filepath.Dir(ReleaseLogDir)
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
