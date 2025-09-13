// Copyright 2024 The lazygophers All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package log 提供了基于 Goroutine 的分布式追踪ID管理功能。
//
// 它通过将追踪ID与 Goroutine ID 绑定，实现了在并发环境下的线程安全追踪。
// 这对于在微服务架构中进行端到端链路追踪尤其有用，可以方便地将一次请求的
// 所有日志串联起来。
//
// # 核心功能
//
//   - 为每个 Goroutine 设置和获取独立的追踪ID。
//   - 自动生成符合规范的追踪ID。
//   - 提供全局开关以禁用追踪功能。
package log

import (
	"crypto/rand"
	"encoding/hex"
	"sync"

	"github.com/petermattis/goid"
)

// traceMap 用于存储 Goroutine ID 到追踪ID的映射。
// 使用 RWMutex + map 组合，在读写性能上优于 sync.Map。
// 键为 int64 类型的 Goroutine ID，值为 string 类型的追踪ID。
var (
	traceMap   = make(map[int64]string, 64) // 预分配容量
	traceMapMu sync.RWMutex
)

// DisableTrace 是一个全局开关，用于禁用或启用追踪功能。
// 当设置为 true 时，所有与追踪ID相关的设置操作（如 SetTrace）将不会生效，
// 获取操作（如 GetTrace）将返回空字符串。
// 这在性能敏感或不需要链路追踪的场景下（如单元测试）非常有用。
var DisableTrace bool

// getTrace 是一个内部函数，用于根据 Goroutine ID 获取其关联的追踪ID。
// 它直接从 traceMap 中加载数据，这是一个线程安全的操作。
//
// 参数:
//
//	gid - Goroutine 的唯一标识符 (int64)。
//
// 返回:
//
//	如果找到了追踪ID，则返回该ID (string)；否则返回一个空字符串。
// 高性能版本：使用内联函数和RLock
//go:inline
func getTrace(gid int64) string {
	traceMapMu.RLock()
	tid := traceMap[gid]
	traceMapMu.RUnlock()
	return tid
}

// setTrace 是一个内部函数，用于为指定的 Goroutine 设置追踪ID。
//
// 如果 traceId 为空字符串，它会自动调用 GenTraceId 生成一个新的、唯一的追踪ID。
// 如果全局开关 DisableTrace 设置为 true，此函数将直接返回，不执行任何操作。
//
// 参数:
//
//	gid     - Goroutine 的唯一标识符 (int64)。
//	traceId - 要设置的追踪ID (string)。如果为空，将自动生成一个新的ID。
//go:inline
func setTrace(gid int64, traceId string) {
	if DisableTrace {
		return
	}
	if traceId == "" {
		traceId = fastGenTraceId()
	}
	traceMapMu.Lock()
	traceMap[gid] = traceId
	traceMapMu.Unlock()
}

// delTrace 是一个内部函数，用于删除指定 Goroutine 的追踪ID。
// 主要用于 Goroutine 生命期结束时清理资源，防止内存泄漏。
//go:inline
func delTrace(gid int64) {
	traceMapMu.Lock()
	delete(traceMap, gid)
	traceMapMu.Unlock()
}

// GetTrace 获取当前 Goroutine 的追踪ID。
//
// 它通过调用 goid.Get() 获取当前 Goroutine 的ID，然后查找并返回与之关联的追踪ID。
// 这是最常用的追踪ID获取方式。
// 如果未设置追踪ID或追踪功能被禁用，将返回空字符串。
//
// 示例:
//
//	id := log.GetTrace()
//	fmt.Printf("Current trace ID: %s\n", id)
func GetTrace() string {
	return getTrace(goid.Get())
}

// GetTraceWithGID 获取指定 Goroutine 的追踪ID。
// 这是一个更底层的函数，允许你查询其他 Goroutine 的追踪ID，但应谨慎使用。
//
// 参数:
//
//	gid - 目标 Goroutine 的唯一标识符 (int64)。
//
// 返回:
//
//	指定 Goroutine 的追踪ID，如果不存在则返回空字符串。
func GetTraceWithGID(gid int64) string {
	return getTrace(gid)
}

// SetTrace 为当前 Goroutine 设置一个追踪ID。
//
// 你可以提供一个或多个字符串作为追踪ID。如果提供了多个，只有第一个会被使用。
// 如果没有提供任何参数 (traceId...)，函数会自动生成一个新的唯一追踪ID。
//
// 示例:
//
//	// 设置一个自定义的追踪ID
//	log.SetTrace("my-custom-trace-id")
//
//	// 自动生成一个新的追踪ID
//	log.SetTrace()
func SetTrace(traceId ...string) {
	// 使用 goid.Get() 获取当前 Goroutine 的 ID
	currentGid := goid.Get()
	if len(traceId) > 0 {
		setTrace(currentGid, traceId[0])
		return
	}
	// 如果没有提供 traceId，则传入空字符串，由 setTrace 内部处理生成逻辑。
	setTrace(currentGid, "")
}

// SetTraceWithGID 为指定的 Goroutine 设置一个追踪ID。
// 允许跨 Goroutine 设置追踪ID，适用于需要手动传递追踪上下文的场景。
//
// 参数:
//
//	gid       - 目标 Goroutine 的唯一标识符 (int64)。
//	traceId... - 一个可选的追踪ID (string)。如果提供多个，仅使用第一个。如果未提供，则自动生成。
func SetTraceWithGID(gid int64, traceId ...string) {
	if len(traceId) > 0 {
		setTrace(gid, traceId[0])
		return
	}
	setTrace(gid, "")
}

// DelTrace 删除当前 Goroutine 的追踪ID。
// 它会清除与当前 Goroutine 关联的追踪ID，使得后续对 GetTrace 的调用返回空字符串。
// 通常在 Goroutine 退出前调用，以释放 traceMap 中的资源。
func DelTrace() {
	delTrace(goid.Get())
}

// DelTraceWithGID 删除指定 Goroutine 的追踪ID。
//
// 参数:
//
//	gid - 目标 Goroutine 的唯一标识符 (int64)。
func DelTraceWithGID(gid int64) {
	delTrace(gid)
}

// fastGenTraceId 生成高性能的追踪ID。
// 使用 crypto/rand 直接生成随机字节，避免 UUID 的额外开销。
//go:inline
func fastGenTraceId() string {
	var buf [8]byte
	rand.Read(buf[:])
	return hex.EncodeToString(buf[:])
}

// GenTraceId 生成一个全局唯一的、16个字符长的追踪ID。
// 保持向后兼容性，内部使用高性能实现。
//
// 返回:
//
//	一个16个字符长的、唯一的追踪ID字符串。
func GenTraceId() string {
	return fastGenTraceId()
}

// 为测试提供的辅助函数
func clearTraceMapForTest() {
	traceMapMu.Lock()
	for k := range traceMap {
		delete(traceMap, k)
	}
	traceMapMu.Unlock()
}

func loadTraceForTest(gid int64) (string, bool) {
	traceMapMu.RLock()
	trace, exists := traceMap[gid]
	traceMapMu.RUnlock()
	return trace, exists
}

