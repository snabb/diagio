package diagio

import (
	"io"
)

// CounterReader implements io.Reader. Count of bytes read is tracked.
type CounterReader struct {
	reader io.Reader
	count  int64
}

var _ io.Reader = (*CounterReader)(nil)

// NewCounterReader wraps io.Reader and returns CounterReader
func NewCounterReader(r io.Reader) (cr *CounterReader) {
	return &CounterReader{
		reader: r,
	}
}

// Read calls Read on the wrapped io.Reader and adds the number of bytes
// read to the counter.
func (cr *CounterReader) Read(p []byte) (n int, err error) {
	n, err = cr.reader.Read(p)
	cr.count = cr.count + int64(n)
	return n, err
}

// Count returns the number of bytes read from the Reader.
func (cr *CounterReader) Count() (n int64) {
	return cr.count
}
