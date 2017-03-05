package diagio

import (
	"encoding/hex"
	"fmt"
	"io"
)

// DiagWriter implements io.Writer. Diagnostics output is produced whenever
// it is written to.
type DiagWriter struct {
	writer    io.Writer
	diagOut   io.Writer
	printInfo bool
	printHex  bool
	pos       int64
}

var _ io.Writer = (*DiagWriter)(nil)

// NewDiagWriter wraps io.Writer and returns DiagWriter.
func NewDiagWriter(w io.Writer, diagOutput io.Writer) (dw *DiagWriter) {
	return &DiagWriter{
		writer:    w,
		diagOut:   diagOutput,
		printInfo: true,
		printHex:  true,
	}
}

// PrintInfo enables or disables output of number of bytes written, stream
// position and possible error code. The default is enabled.
func (dw *DiagWriter) PrintInfo(v bool) {
	dw.printInfo = v
}

// PrintHex enables or disables hex dump output.
// The default is enabled.
func (dw *DiagWriter) PrintHex(v bool) {
	dw.printHex = v
}

// Write calls Write on the wrapped io.Writer and produces diagnostics output.
func (dw *DiagWriter) Write(p []byte) (n int, err error) {
	if dw.printInfo {
		fmt.Fprintf(dw.diagOut, ">>> Write %d bytes at %d (start) ===\n", len(p), dw.pos)
	}
	if dw.printHex && len(p) > 0 {
		fmt.Fprint(dw.diagOut, hex.Dump(p))
	}

	n, err = dw.writer.Write(p)

	if dw.printInfo {
		fmt.Fprintf(dw.diagOut, "<<< Wrote %d bytes (err = %v) (end) ===\n", n, err)
	}
	dw.pos = dw.pos + int64(n)

	return n, err
}
