package log

import (
	"bytes"
	"sync"
)

// bufPool a pool of byte buffers.
//
// It is used to reuse pre-allocated buffers, reducing memory allocation overhead.
//
// How sync.Pool works:
//   - Each processor (P) maintains a local pool, reducing lock contention.
//   - When the local pool is empty, it will try to steal from other processors or create a new object.
//   - Objects in the pool may be cleaned up during GC.
//
// Use cases:
//   - Suitable for objects that are frequently created/destroyed and have a high creation cost.
//   - Especially suitable for reusing buffers in high-frequency operations such as logging.
var bufPool = sync.Pool{
	New: func() any {
		// Create a new bytes.Buffer instance.
		return &bytes.Buffer{}
	},
}

// GetBuffer retrieves a buffer from the pool.
//
// It returns a reusable *bytes.Buffer instance.
func GetBuffer() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

// PutBuffer returns a buffer to the pool.
//
// The buffer is reset before being put back into the pool.
// If the buffer is nil, the function will return directly.
//
// buf is the buffer instance to be returned.
func PutBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	buf.Reset()
	bufPool.Put(buf)
}
