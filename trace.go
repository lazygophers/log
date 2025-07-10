// Package log 提供分布式追踪ID管理功能
//
// 使用goroutine本地存储实现线程安全的追踪ID管理，适用于微服务链路追踪场景
package log

import (
	"strings"
	"sync"

	"github.com/google/uuid"

	"github.com/petermattis/goid"
)

// traceMap 存储goroutine ID到追踪ID的映射
// 使用sync.Map实现线程安全访问
var traceMap sync.Map

// DisableTrace 全局追踪开关
// 设置为true时禁用所有追踪操作
var DisableTrace bool

// getTrace 获取指定goroutine的追踪ID
// 参数: gid - goroutine ID
// 返回: 追踪ID字符串，不存在时返回空字符串
func getTrace(gid int64) string {
	tid, ok := traceMap.Load(gid)
	if !ok {
		return ""
	}
	return tid.(string)
}

// setTrace 设置指定goroutine的追踪ID
// 参数:
//
//	gid - goroutine ID
//	traceId - 要设置的追踪ID（为空时自动生成）
//
// 注意: 当DisableTrace=true时此操作无效
func setTrace(gid int64, traceId string) {
	if DisableTrace {
		return
	}
	if traceId == "" {
		traceId = GenTraceId()
	}
	traceMap.Store(gid, traceId)
}

// delTrace 删除指定goroutine的追踪ID
// 参数: gid - goroutine ID
func delTrace(gid int64) {
	traceMap.Delete(gid)
}

// GetTrace 获取当前goroutine的追踪ID
// 返回: 当前goroutine的追踪ID字符串
func GetTrace() string {
	return getTrace(goid.Get())
}

// GetTraceWithGID 获取指定goroutine的追踪ID
// 参数: gid - 目标goroutine ID
// 返回: 指定goroutine的追踪ID字符串
func GetTraceWithGID(gid int64) string {
	return getTrace(gid)
}

// SetTrace 设置当前goroutine的追踪ID
// 参数: traceId... - 可选追踪ID（不提供则自动生成）
func SetTrace(traceId ...string) {
	if len(traceId) > 0 {
		setTrace(goid.Get(), traceId[0])
		return
	}
	setTrace(goid.Get(), "")
}

// SetTraceWithGID 设置指定goroutine的追踪ID
// 参数:
//
//	gid - 目标goroutine ID
//	traceId... - 可选追踪ID（不提供则自动生成）
func SetTraceWithGID(gid int64, traceId ...string) {
	if len(traceId) > 0 {
		setTrace(gid, traceId[0])
		return
	}
	setTrace(gid, "")
}

// DelTrace 删除当前goroutine的追踪ID
func DelTrace() {
	delTrace(goid.Get())
}

// DelTraceWithGID 删除指定goroutine的追踪ID
// 参数: gid - 目标goroutine ID
func DelTraceWithGID(gid int64) {
	delTrace(gid)
}

// GenTraceId 生成追踪ID
// 算法: 基于UUIDv4，移除连字符后取后16位字符
// 返回: 16字符长度的追踪ID字符串
func GenTraceId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")[16:]
}
