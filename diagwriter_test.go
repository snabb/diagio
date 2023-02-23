package diagio_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/snabb/diagio"
)

func ExampleDiagWriter() {
	dw := diagio.NewDiagWriter(io.Discard, os.Stdout)
	fmt.Fprint(dw, "hello world")
	// Output:
	// >>> Write 11 bytes at 0 (start) ===
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
	// <<< Wrote 11 bytes (err = <nil>) (end) ===
}

func TestDiagWriter(t *testing.T) {
	output := new(bytes.Buffer)
	dw := diagio.NewDiagWriter(io.Discard, output)
	dw.PrintHex(false)
	dw.PrintInfo(false)
	fmt.Fprint(dw, "foobar")
	if output.Len() != 0 {
		t.Errorf("Emitted %d bytes when both outputs were turned off", output.Len())
	}

	output.Reset()
	dw.PrintHex(true)
	dw.PrintInfo(false)
	fmt.Fprint(dw, "foobar")
	if output.Len() == 0 {
		t.Error("No output produced when hex output was enabled")
	}

	output.Reset()
	dw.PrintHex(false)
	dw.PrintInfo(true)
	fmt.Fprint(dw, "foobar")
	if output.Len() == 0 {
		t.Error("No output produced when info output was enabled")
	}
}
