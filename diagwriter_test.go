package diagio_test

import (
	"fmt"
	"github.com/snabb/diagio"
	"io/ioutil"
	"os"
)

func ExampleDiagWriter() {
	output := diagio.NewDiagWriter(ioutil.Discard, os.Stdout)
	fmt.Fprint(output, "hello world")
	// Output:
	// >>> Write 11 bytes at 0 (start) ===
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
	// <<< Wrote 11 bytes (err = <nil>) (end) ===
}
