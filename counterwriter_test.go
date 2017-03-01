package diagio_test

import (
	"fmt"
	"github.com/snabb/diagio"
	"os"
)

func ExampleWriter() {
	cw := diagio.NewCounterWriter(os.Stdout)
	fmt.Fprintln(cw, "hello world")
	fmt.Println(cw.Count())
	// Output:
	// hello world
	// 12
}
