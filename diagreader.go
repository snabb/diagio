package diagio

import (
	"encoding/hex"
	"fmt"
	"io"
)

// DiagReader implements io.Reader. Diagnostics output is produced whenever
// it is read from.
type DiagReader struct {
	reader    io.Reader
	diagOut   io.Writer
	printInfo bool
	printHex  bool
	pos       int64
}

var _ io.Reader = (*DiagReader)(nil)

// NewDiagReader wraps io.Reader and returns DiagReader.
func NewDiagReader(r io.Reader, diagOutput io.Writer) (dr *DiagReader) {
	return &DiagReader{
		reader:    r,
		diagOut:   diagOutput,
		printInfo: true,
		printHex:  true,
	}
}

// PrintInfo enables or disables output of number of bytes read, stream
// position and possible error code. The default is enabled.
func (dr *DiagReader) PrintInfo(v bool) {
	dr.printInfo = v
}

// PrintHex enables or disables hex dump output.
// The default is enabled.
func (dr *DiagReader) PrintHex(v bool) {
	dr.printHex = v
}

// Read calls Read on the wrapped io.Reader and produces diagnostics output.
func (dr *DiagReader) Read(p []byte) (n int, err error) {
	if dr.printInfo {
		fmt.Fprintf(dr.diagOut, ">>> Read %d bytes at %d (start) ===\n", len(p), dr.pos)
	}

	n, err = dr.reader.Read(p)

	if dr.printHex && n > 0 {
		fmt.Fprint(dr.diagOut, hex.Dump(p[0:n]))
	}
	if dr.printInfo {
		fmt.Fprintf(dr.diagOut, "<<< Read %d bytes (err = %v) (end) ===\n", n, err)
	}
	dr.pos = dr.pos + int64(n)

	return n, err
}
