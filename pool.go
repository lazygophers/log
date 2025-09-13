package log

import (
	"bytes"
	"sync"
)

// bufPool is a byte buffer pool for reusing pre-allocated buffers
var bufPool = sync.Pool{
	New: func() any {
		// Create a new bytes.Buffer instance
		return &bytes.Buffer{}
	},
}

// GetBuffer retrieves a buffer from the pool
func GetBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

// PutBuffer returns a buffer to the pool after resetting it
func PutBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	buf.Reset()
	bufPool.Put(buf)
}
