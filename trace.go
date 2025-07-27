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
	"strings"
	"sync"

	"github.com/google/uuid"

	"github.com/petermattis/goid"
)

// traceMap 用于存储 Goroutine ID 到追踪ID的映射。
// 它是一个线程安全的 map，键为 int64 类型的 Goroutine ID，值为 string 类型的追踪ID。
var traceMap sync.Map

// DisableTrace 是一个全局开关，用于禁用或启用追踪功能。
// 当设置为 true 时，所有与追踪ID相关的设置操作（如 SetTrace）将不会生效，
// 获取操作（如 GetTrace）将返回空字符串。这在测试或不需要追踪的场景下很有用。
var DisableTrace bool

// getTrace 是一个内部函数，用于获取指定 Goroutine 的追踪ID。
// 它直接从 traceMap 中加载与给定 gid 关联的追踪ID。
//
// 参数:
//
//	gid - Goroutine 的唯一标识符。
//
// 返回:
//
//	如果找到了追踪ID，则返回该ID；否则返回一个空字符串。
func getTrace(gid int64) string {
	tid, ok := traceMap.Load(gid)
	if !ok {
		return ""
	}
	return tid.(string)
}

// setTrace 是一个内部函数，用于设置指定 Goroutine 的追踪ID。
//
// 如果 traceId 为空字符串，它会自动生成一个新的追踪ID。
// 如果 DisableTrace 设置为 true，此函数将不执行任何操作。
//
// 参数:
//
//	gid     - Goroutine 的唯一标识符。
//	traceId - 要设置的追踪ID。如果为空，将自动生成一个新的ID。
func setTrace(gid int64, traceId string) {
	if DisableTrace {
		return
	}
	if traceId == "" {
		traceId = GenTraceId()
	}
	traceMap.Store(gid, traceId)
}

// delTrace 是一个内部函数，用于删除指定 Goroutine 的追踪ID。
func delTrace(gid int64) {
	traceMap.Delete(gid)
}

// GetTrace 获取当前 Goroutine 的追踪ID。
//
// 它会自动获取当前 Goroutine 的 ID，并返回与之关联的追踪ID。
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
//
// 参数:
//
//	gid - 目标 Goroutine 的唯一标识符。
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
// 如果没有提供任何参数，函数会自动生成一个新的唯一追踪ID。
//
// 示例:
//
//	// 设置一个自定义的追踪ID
//	log.SetTrace("my-custom-trace-id")
//
//	// 自动生成一个新的追踪ID
//	log.SetTrace()
func SetTrace(traceId ...string) {
	if len(traceId) > 0 {
		setTrace(goid.Get(), traceId[0])
		return
	}
	setTrace(goid.Get(), "")
}

// SetTraceWithGID 为指定的 Goroutine 设置一个追踪ID。
//
// 参数:
//
//	gid       - 目标 Goroutine 的唯一标识符。
//	traceId... - 一个可选的追踪ID。如果提供多个，仅使用第一个。如果未提供，则自动生成。
func SetTraceWithGID(gid int64, traceId ...string) {
	if len(traceId) > 0 {
		setTrace(gid, traceId[0])
		return
	}
	setTrace(gid, "")
}

// DelTrace 删除当前 Goroutine 的追踪ID。
// 它会清除与当前 Goroutine 关联的追踪ID，使得后续对 GetTrace 的调用返回空字符串。
func DelTrace() {
	delTrace(goid.Get())
}

// DelTraceWithGID 删除指定 Goroutine 的追踪ID。
//
// 参数:
//
//	gid - 目标 Goroutine 的唯一标识符。
func DelTraceWithGID(gid int64) {
	delTrace(gid)
}

// GenTraceId 生成一个全局唯一的追踪ID。
//
// 算法:
//  1. 生成一个 UUID v4。
//  2. 移除其中的连字符 '-'。
//  3. 截取最后16个字符作为最终的追踪ID。
//
// 返回:
//
//	一个16个字符长的、唯一的追踪ID字符串。
func GenTraceId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")[16:]
}
