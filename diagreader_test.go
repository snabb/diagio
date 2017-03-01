package diagio_test

import (
	"bytes"
	"github.com/snabb/diagio"
	"io"
	"io/ioutil"
	"os"
)

func ExampleDiagReader() {
	input1 := bytes.NewBufferString("hello world")
	input2 := diagio.NewDiagReader(input1, os.Stdout)
	io.Copy(ioutil.Discard, input2)
	// Output:
	// === Read 11 bytes at 0 (err = <nil>) ===
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
	// === Read 0 bytes at 11 (err = EOF) ===
}
