package log

import (
	"testing"
)

func TestFormatFull(t *testing.T) {
	// 测试formatter.format()和Format()的逻辑覆盖
	f := &Formatter{}
	entry := &Entry{}
	result := f.format(entry)
	if len(result) == 0 {
		t.Fatal("format should not return empty")
	}
	// 断言结果是否符合预期
}

func TestGetColorByLevel(t *testing.T) {
	// 测试不同level对应的颜色输出
	// 覆盖所有Level枚举值
}
