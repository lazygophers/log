package log

import (
	"bytes"
	"errors"
	"sync"
)

type AsyncWriter struct {
	writer Writer

	c chan []byte

	close chan *sync.WaitGroup
}

var ErrAsyncWriterFull = errors.New("async writer full")

func (p *AsyncWriter) Write(b []byte) (n int, err error) {
	select {
	case p.c <- b:
		return len(b), nil
	default:
		return 0, ErrAsyncWriterFull
	}
}

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

func NewAsyncWriter(writer Writer) *AsyncWriter {
	p := &AsyncWriter{
		writer: writer,
		c:      make(chan []byte, 1024),
		close:  make(chan *sync.WaitGroup, 1),
	}

	go func() {
		var cache bytes.Buffer
		for {
			cache.Reset()
			select {
			case c := <-p.c:
				cache.Write(c)
				for {
					select {
					case c := <-p.c:
						cache.Write(c)
					default:
						goto OUT
					}
				}
			OUT:
				p.writer.Write(cache.Bytes())

			case w := <-p.close:
				for {
					select {
					case c := <-p.c:
						cache.Write(c)
					default:
						goto END
					}
				}
			END:
				p.writer.Write(cache.Bytes())
				w.Done()

				return
			}
		}
	}()

	return p
}
