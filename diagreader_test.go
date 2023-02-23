package diagio_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/snabb/diagio"
)

func ExampleDiagReader() {
	input := bytes.NewBufferString("hello world")
	dr := diagio.NewDiagReader(input, os.Stdout)
	io.Copy(io.Discard, dr)
	// Output:
	// >>> Read 8192 bytes at 0 (start) ===
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
	// <<< Read 11 bytes (err = <nil>) (end) ===
	// >>> Read 8192 bytes at 11 (start) ===
	// <<< Read 0 bytes (err = EOF) (end) ===
}

func TestDiagReader(t *testing.T) {
	input := bytes.NewBufferString("foobar")
	output := new(bytes.Buffer)
	dr := diagio.NewDiagReader(input, output)
	dr.PrintHex(false)
	dr.PrintInfo(false)
	io.Copy(io.Discard, dr)
	if output.Len() != 0 {
		t.Errorf("Emitted %d bytes when both outputs were turned off", output.Len())
	}

	input = bytes.NewBufferString("foobar")
	output.Reset()
	dr = diagio.NewDiagReader(input, output)
	dr.PrintHex(true)
	dr.PrintInfo(false)
	io.Copy(io.Discard, dr)
	if output.Len() == 0 {
		t.Error("No output produced when hex output was enabled")
	}

	input = bytes.NewBufferString("foobar")
	output.Reset()
	dr = diagio.NewDiagReader(input, output)
	dr.PrintHex(false)
	dr.PrintInfo(true)
	io.Copy(io.Discard, dr)
	if output.Len() == 0 {
		t.Error("No output produced when info output was enabled")
	}
}
