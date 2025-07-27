// Package log 提供异步日志写入器实现
// 通过通道缓冲和批量写入优化高并发场景下的日志写入性能。
package log

import (
	"bytes"
	"errors"
	"sync"
)

// AsyncWriter 定义了一个异步日志写入器。
// 它通过一个缓冲通道接收日志数据，并由一个后台 goroutine 批量写入底层 writer，
// 从而实现非阻塞的日志写入操作，在高并发场景下可以显著提升性能。
type AsyncWriter struct {
	writer Writer               // writer 是实际执行写入操作的底层写入器。
	c      chan []byte          // c 是用于缓冲日志数据的通道。
	close  chan *sync.WaitGroup // close 是用于接收关闭信号的通道，并使用 WaitGroup 实现同步关闭。
}

// ErrAsyncWriterFull 异步写入器通道满错误。
// 当缓冲通道满时写入会返回此错误。
var ErrAsyncWriterFull = errors.New("async writer full")

// Write 将字节数据 b 异步写入。
// 它将数据发送到缓冲通道 c，如果通道已满，则立即返回 ErrAsyncWriterFull 错误，
// 实现了非阻塞写入。实际的写入操作由后台的 goroutine 完成。
func (p *AsyncWriter) Write(b []byte) (n int, err error) {
	select {
	case p.c <- b:
		return len(b), nil
	default:
		return 0, ErrAsyncWriterFull
	}
}

// Close 关闭异步写入器。
// 这是一个优雅关闭的实现：它会向 close 通道发送一个关闭信号，
// 并等待后台 goroutine 将缓冲通道 c 中的所有剩余数据都写入底层 writer 后再返回。
// 此方法是线程安全的，可以被多次调用。
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

// NewAsyncWriter 创建并初始化一个 AsyncWriter 实例。
// 它接收一个底层的 Writer 接口作为参数，并启动一个后台 goroutine 来处理日志的批量写入。
//
// 参数:
//
//	writer: 一个实现了 Writer 接口的实例，用于最终的日志输出。
//
// 返回:
//
//	一个初始化完成的 *AsyncWriter 指针。
func NewAsyncWriter(writer Writer) *AsyncWriter {
	p := &AsyncWriter{
		writer: writer,
		c:      make(chan []byte, 1024),       // 初始化容量为 1024 的缓冲通道
		close:  make(chan *sync.WaitGroup, 1), // 初始化容量为 1 的关闭信号通道
	}

	// 启动一个后台 goroutine 负责消费通道中的数据并批量写入。
	go func() {
		var cache bytes.Buffer // 创建一个字节缓冲区，用于汇集多条日志
		for {
			cache.Reset() // 每次循环开始时重置缓冲区
			select {
			// case1: 正常接收日志数据
			case c := <-p.c:
				// 将第一条日志写入缓冲区
				cache.Write(c)
				// 通过内层循环，尽可能多地从通道中读取数据，实现批量处理
				for {
					select {
					case c := <-p.c:
						cache.Write(c) // 将后续日志追加到缓冲区
					default:
						goto OUT // 如果通道为空，则跳出内层循环，执行写入
					}
				}
			OUT:
				// 将汇集的多条日志一次性写入底层 writer
				p.writer.Write(cache.Bytes())

			// case2: 接收到关闭信号，准备优雅退出
			case w := <-p.close:
				// 在退出前，需要处理通道中所有剩余的日志
				for {
					select {
					case c := <-p.c:
						cache.Write(c) // 将剩余日志追加到缓冲区
					default:
						goto END // 如果通道已空，则跳出循环
					}
				}
			END:
				// 将缓冲区中最后剩余的日志写入底层 writer
				if cache.Len() > 0 {
					p.writer.Write(cache.Bytes())
				}
				w.Done() // 通知 Close() 方法，清理工作已完成
				return   // 退出 goroutine
			}
		}
	}()

	return p
}
