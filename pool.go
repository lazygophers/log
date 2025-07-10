package log

import (
	"bytes"
	"sync"
)

// bufPool 是字节缓冲区的对象池，用于复用已分配的缓冲区，减少内存分配开销。
// sync.Pool 实现原理：
//   - 每个处理器(P)维护一个本地池，减少锁竞争
//   - 当本地池为空时，会尝试从其他处理器偷取或创建新对象
//   - 池中对象可能在GC时被清理
//
// 使用场景：
//   - 适用于频繁创建/销毁且创建成本高的对象
//   - 特别适合在日志记录等高频操作中复用缓冲区
var bufPool = sync.Pool{
	New: func() interface{} {
		// 创建新的bytes.Buffer实例
		return &bytes.Buffer{}
	},
}

// GetBuffer 从对象池中获取一个字节缓冲区。
// 返回值:
//
//	*bytes.Buffer - 可重用的缓冲区实例
func GetBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

// PutBuffer 将使用完毕的缓冲区归还到对象池，以便复用。
// 参数:
//
//	buf *bytes.Buffer - 需要归还的缓冲区实例
//
// 注意:
//   - 如果传入nil，函数将直接返回
//   - 缓冲区在归还前会被重置(Reset)
func PutBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	buf.Reset()
	bufPool.Put(buf)
}
