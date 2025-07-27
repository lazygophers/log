package log

import "io"

// Writer 定义了一个可关闭的写入器 (Writer) 接口。
// 它聚合了标准的 io.Writer 接口，并额外增加了一个 Close 方法。
type Writer interface {
	// 继承标准库的 io.Writer 接口，使其具备写入数据的能力。
	io.Writer

	// Close 用于关闭写入器，并释放所有占用的资源。
	// 在写入操作完成后，应始终调用此方法。
	// 如果关闭过程中发生错误，将返回一个非空的 error。
	Close() error
}
