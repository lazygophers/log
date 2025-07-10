package log

import "io"

// Writer 定义了带关闭功能的写入器接口
type Writer interface {
	io.Writer

	// Close 关闭写入器并释放资源
	// 返回可能出现的错误
	Close() error
}
