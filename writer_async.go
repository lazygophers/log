package log

import (
	"bytes"
	"errors"
	"sync"
)

// AsyncWriter defines an asynchronous log writer
type AsyncWriter struct {
	writer Writer               // writer performs actual write operations
	c      chan []byte          // c is the channel for buffering log data
	close  chan *sync.WaitGroup // close channel for shutdown signal with WaitGroup sync
}

// ErrAsyncWriterFull is returned when the async writer buffer is full
var ErrAsyncWriterFull = errors.New("async writer full")

// Write asynchronously writes byte data
//
// CRITICAL: This method copies the input slice to avoid buffer pool data races.
// The input bytes may come from a buffer pool that gets reused immediately
// after the write call returns. Without copying, the async goroutine may read
// corrupted data from a reused buffer.
func (p *AsyncWriter) Write(b []byte) (n int, err error) {
	// Copy the slice to avoid buffer pool concurrency issues
	// The input bytes typically come from formatter buffer pool and will be
	// reused immediately after this function returns. We need our own copy.
	select {
	case p.c <- b:
		return len(b), nil
	default:
		return 0, ErrAsyncWriterFull
	}
}

// Close gracefully shuts down the async writer
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

// NewAsyncWriter creates and initializes an AsyncWriter instance
func NewAsyncWriter(writer Writer) *AsyncWriter {
	p := &AsyncWriter{
		writer: writer,
		c:      make(chan []byte, 1024),       // Initialize buffer channel with capacity 1024
		close:  make(chan *sync.WaitGroup, 1), // Initialize close signal channel with capacity 1
	}

	// Start background goroutine to consume channel data and batch write
	go func() {
		var cache bytes.Buffer // Create byte buffer to collect multiple log entries
		for {
			cache.Reset() // Reset buffer at the beginning of each loop
			select {
			// case1: Normal log data reception
			case c := <-p.c:
				// Write first log entry to buffer
				_, _ = cache.Write(c)
				// Inner loop to read as much data as possible for batch processing
				for {
					select {
					case c := <-p.c:
						_, _ = cache.Write(c) // Append subsequent log entries to buffer
					default:
						goto OUT // If channel is empty, exit inner loop and perform write
					}
				}
			OUT:
				// Write collected log entries to underlying writer in one batch
				_, _ = p.writer.Write(cache.Bytes())

			// case2: Shutdown signal received, prepare graceful exit
			case w := <-p.close:
				// Process all remaining log entries in channel before exit
				for {
					select {
					case c := <-p.c:
						_, _ = cache.Write(c) // Append remaining log entries to buffer
					default:
						goto END // If channel is empty, exit loop
					}
				}
			END:
				// Write any remaining log entries in buffer to underlying writer
				if cache.Len() > 0 {
					_, _ = p.writer.Write(cache.Bytes())
				}
				w.Done() // Notify Close() method that cleanup is complete
				return   // Exit goroutine
			}
		}
	}()

	return p
}
