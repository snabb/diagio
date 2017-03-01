package diagio_test

import (
	"fmt"
	"github.com/snabb/diagio"
	"os"
)

func ExampleCounterWriter() {
	cw := diagio.NewCounterWriter(os.Stdout)
	fmt.Fprintln(cw, "hello world")
	fmt.Println(cw.Count())
	// Output:
	// hello world
	// 12
}
