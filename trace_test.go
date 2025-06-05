package log

import (
	"testing"

	"github.com/petermattis/goid"
	"github.com/stretchr/testify/assert"
)

// TestGenTraceId 验证生成的traceID格式
func TestGenTraceId(t *testing.T) {
	id := GenTraceId()
	assert.Len(t, id, 16, "traceID应该恰好16字符")
	assert.NotEmpty(t, id, "traceID不能为空")
}

// TestSetTrace 验证traceID设置逻辑
func TestSetTrace(t *testing.T) {
	gid := goid.Get()

	// 基础设置
	SetTrace("base-id")
	assert.Equal(t, "base-id", GetTrace(), "应该获取到设置的traceID")

	// GID设置
	SetTraceWithGID(gid, "custom-id")
	assert.Equal(t, "custom-id", GetTraceWithGID(gid), "通过GID获取应该返回指定ID")

	// 空ID自动生成
	SetTrace("")
	assert.Len(t, GetTrace(), 16, "空ID应该自动生成16字符")
}

// TestDelTrace 验证traceID删除逻辑
func TestDelTrace(t *testing.T) {
	gid := goid.Get()

	SetTrace("to-delete")
	assert.NotEmpty(t, GetTrace(), "设置ID前应该存在")

	DelTrace()
	assert.Empty(t, GetTrace(), "删除后应该为空")

	SetTraceWithGID(gid, "another-id")
	assert.NotEmpty(t, GetTraceWithGID(gid), "通过GID设置后应该存在")

	DelTraceWithGID(gid)
	assert.Empty(t, GetTraceWithGID(gid), "通过GID删除后应该为空")
}

// TestGetTrace 验证获取traceID的准确性
func TestGetTrace(t *testing.T) {
	gid := goid.Get()
	SetTrace("direct-id")
	SetTraceWithGID(gid, "explicit-id")

	assert.Equal(t, "explicit-id", GetTraceWithGID(gid), "GID获取应该返回显式设置的ID")
	assert.Equal(t, "explicit-id", GetTrace(), "默认获取应该返回相同ID")
}

// TestTraceDisable 验证禁用追踪功能
func TestTraceDisable(t *testing.T) {
	oldState := DisableTrance
	defer func() { DisableTrance = oldState }()

	DisableTrance = true
	SetTrace("should-fail")
	assert.Empty(t, GetTrace(), "禁用追踪时应该不存储ID")

	DisableTrance = false
	SetTrace("should-work")
	assert.Equal(t, "should-work", GetTrace(), "启用追踪时应该存储ID")
}
