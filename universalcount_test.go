package diagio_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/snabb/diagio"
)

func ExampleCount() {
	cw := diagio.NewCounterWriter(os.Stdout)

	f := func(w io.Writer) {
		fmt.Fprintln(w, "foobar")
		fmt.Println(diagio.Count(w))
	}

	f(cw)
	// Output:
	// foobar
	// 7
}

func TestCount(t *testing.T) {
	n := diagio.Count(nil)
	if n != -1 {
		t.Errorf("Count return value %d != -1", n)
	}
}
