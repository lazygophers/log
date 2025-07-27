package log

import (
	"testing"
)

func TestGetBuffer(t *testing.T) {
	buf := GetBuffer()
	if buf == nil {
		t.Fatal("expected to get a buffer, but got nil")
	}
	if buf.Len() != 0 {
		t.Errorf("expected buffer length 0, but got %d", buf.Len())
	}
	if buf.Cap() != 0 {
		t.Errorf("expected buffer capacity 0, but got %d", buf.Cap())
	}
}

func TestPutBuffer(t *testing.T) {

}

func BenchmarkPool(b *testing.B) {

}
