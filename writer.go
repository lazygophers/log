package log

import "io"

type Writer interface {
	io.Writer
	Close() error
}
