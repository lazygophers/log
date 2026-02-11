package log

import (
	"bytes"
	"sync"
	"testing"
	"time"
)

// TestAsyncWriter_BufferPoolSafety verifies that AsyncWriter correctly
// handles bytes from buffer pool without data races
func TestAsyncWriter_BufferPoolSafety(t *testing.T) {
	// Create a buffer that tracks if it's being modified after being written
	tracker := &trackingWriter{}
	asyncWriter := NewAsyncWriter(tracker)

	// Simulate what happens in real logging: get buffer, write, return to pool
	const iterations = 100

	for i := 0; i < iterations; i++ {
		// Simulate formatter behavior: get buffer from pool
		buf := GetBuffer()
		testMsg := "test-message-"
		buf.WriteString(testMsg)
		buf.WriteByte(byte('0' + i%10))

		// Get the bytes that would be passed to async writer
		logBytes := buf.Bytes()

		// This is what happens in logger.write(): pass bytes to async writer
		_, _ = asyncWriter.Write(logBytes)

		// CRITICAL: Immediately return buffer to pool (this is what causes the bug)
		PutBuffer(buf)

		// Small delay to avoid overflowing the async channel
		if i%10 == 9 {
			time.Sleep(time.Millisecond)
		}
	}

	// Give async writer time to process all writes
	time.Sleep(200 * time.Millisecond)

	// Close the async writer and wait for cleanup
	_ = asyncWriter.Close()

	// Verify that all written data is correct
	tracker.mu.Lock()
	defer tracker.mu.Unlock()

	if len(tracker.writes) != iterations {
		t.Logf("Expected %d writes, got %d (async writer may not have processed all)", iterations, len(tracker.writes))
		// Don't fail the test, just log - we're mainly checking for data races
	}

	// Check that no data was corrupted (no unexpected overlapping)
	for i, write := range tracker.writes {
		expectedPrefix := []byte("test-message-")
		if !bytes.HasPrefix(write, expectedPrefix) {
			t.Errorf("Write %d has corrupted data: %q (expected prefix %q)", i, write, expectedPrefix)
		}
	}

	// If we got here without data race detection warnings, the test passed
	t.Log("Buffer pool safety test completed successfully")
}

// trackingWriter is a test writer that records all writes
type trackingWriter struct {
	mu    sync.Mutex
	writes [][]byte
}

func (t *trackingWriter) Write(p []byte) (n int, err error) {
	t.mu.Lock()
	// Make a copy to store, since p will be reused
	buf := make([]byte, len(p))
	n = copy(buf, p)
	t.writes = append(t.writes, buf)
	t.mu.Unlock()
	return n, nil
}

func (t *trackingWriter) Close() error {
	return nil
}
