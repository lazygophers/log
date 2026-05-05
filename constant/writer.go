package constant

import (
	"io"
)

// AddSync converts io.Writer to WriteSyncer
func AddSync(w io.Writer) WriteSyncer {
	if ws, ok := w.(WriteSyncer); ok {
		return ws
	}
	return &writeSyncerAdapter{Writer: w}
}

type writeSyncerAdapter struct {
	io.Writer
}

func (a *writeSyncerAdapter) Sync() error {
	return nil
}

// NewMultiWriteSyncer creates a writer that writes to multiple writers
func NewMultiWriteSyncer(writers ...WriteSyncer) WriteSyncer {
	return &multiWriteSyncer{writers: writers}
}

type multiWriteSyncer struct {
	writers []WriteSyncer
}

func (m *multiWriteSyncer) Write(p []byte) (n int, err error) {
	for _, w := range m.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (m *multiWriteSyncer) Sync() error {
	for _, w := range m.writers {
		if err := w.Sync(); err != nil {
			return err
		}
	}
	return nil
}
