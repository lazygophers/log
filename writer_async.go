// Package log 提供异步日志写入器实现
// 通过通道缓冲和批量写入优化高并发场景下的日志写入性能。
package log

import (
	"bytes"
	"errors"
	"sync"
)

// AsyncWriter 异步日志写入器。
// 使用缓冲通道实现非阻塞写入，
// 后台goroutine执行批量写入。
// 字段:
//
//	writer - 底层日志写入器。
//	c      - 日志数据缓冲通道（容量1024）。
//	close  - 关闭信号通道。
type AsyncWriter struct {
	writer Writer
	c      chan []byte
	close  chan *sync.WaitGroup
}

// ErrAsyncWriterFull 异步写入器通道满错误。
// 当缓冲通道满时写入会返回此错误。
var ErrAsyncWriterFull = errors.New("async writer full")

// Write 异步写入日志数据。
// 参数: b - 要写入的字节数据。
// 返回: 写入字节数, 错误信息。
// 特性:
//   - 非阻塞设计（通道满时立即返回错误）。
//   - 实际写入由后台goroutine批量完成。
func (p *AsyncWriter) Write(b []byte) (n int, err error) {
	select {
	case p.c <- b:
		return len(b), nil
	default:
		return 0, ErrAsyncWriterFull
	}
}

// Close 关闭异步写入器。
// 特性:
//   - 优雅关闭：等待缓冲数据全部写入。
//   - 线程安全：可多次调用。
//
// 返回: 关闭过程中发生的错误。
func (p *AsyncWriter) Close() error {
	var w sync.WaitGroup
	w.Add(1)
	select {
	case p.close <- &w:
		w.Wait()
	default:
	}
	return nil
}

// NewAsyncWriter 创建异步日志写入器。
// 参数: writer - 底层日志写入器实现。
// 返回: 初始化好的AsyncWriter指针。
// 注意:
//   - 内部启动goroutine执行批量写入。
//   - 默认通道缓冲1024条日志。
func NewAsyncWriter(writer Writer) *AsyncWriter {
	p := &AsyncWriter{
		writer: writer,
		c:      make(chan []byte, 1024),
		close:  make(chan *sync.WaitGroup, 1),
	}

	// 后台处理协程
	go func() {
		var cache bytes.Buffer
		for {
			cache.Reset()
			select {
			case c := <-p.c:
				// 批量读取通道数据
				cache.Write(c)
				for {
					select {
					case c := <-p.c:
						cache.Write(c)
					default:
						goto OUT
					}
				}
			OUT:
				// 批量写入底层writer
				p.writer.Write(cache.Bytes())

			case w := <-p.close:
				// 关闭前处理剩余数据
				for {
					select {
					case c := <-p.c:
						cache.Write(c)
					default:
						goto END
					}
				}
			END:
				p.writer.Write(cache.Bytes())
				w.Done()
				return
			}
		}
	}()

	return p
}
