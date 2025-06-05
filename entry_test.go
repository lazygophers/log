package log

import (
	"testing"
)

func TestEntryReset(t *testing.T) {
	// 测试Entry.Reset()的逻辑覆盖
	e := &Entry{}
	e.Reset()
	// 断言字段是否重置
}
