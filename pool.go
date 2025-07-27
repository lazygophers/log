package log

import (
	"bytes"
	"sync"
)

// bufPool 是一个字节缓冲区的池。
// 它用于重用预先分配的缓冲区，从而减少内存分配开销。
//
// sync.Pool 的工作原理：
//   - 每个处理器（P）都维护一个本地池，以减少锁竞争。
//   - 当本地池为空时，它会尝试从其他处理器窃取或创建一个新对象。
//   - 池中的对象可能会在垃圾回收（GC）期间被清理。
//
// 使用场景：
//   - 适用于频繁创建/销毁且创建成本较高的对象。
//   - 特别适用于在日志记录等高频操作中重用缓冲区。
var bufPool = sync.Pool{
	New: func() any {
		// 创建一个新的 bytes.Buffer 实例。
		return &bytes.Buffer{}
	},
}

// GetBuffer 从池中检索一个缓冲区。
// 它返回一个可重用的 *bytes.Buffer 实例。
func GetBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

// PutBuffer 将缓冲区返回到池中。
// 在放回池中之前，缓冲区会被重置。
// 如果缓冲区为 nil，函数将直接返回。
//
// buf 是要返回的缓冲区实例。
func PutBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	buf.Reset()
	bufPool.Put(buf)
}
