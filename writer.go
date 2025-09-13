package log

import "io"

// Writer defines a closable writer interface
type Writer interface {
	// Inherits standard io.Writer interface for data writing capability
	io.Writer

	// Close closes the writer and releases all resources
	Close() error
}
