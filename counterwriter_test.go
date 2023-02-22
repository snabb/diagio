package diagio_test

import (
	"fmt"
	"os"

	"github.com/snabb/diagio"
)

func ExampleCounterWriter() {
	cw := diagio.NewCounterWriter(os.Stdout)
	fmt.Fprintln(cw, "hello world")
	fmt.Println(cw.Count())
	// Output:
	// hello world
	// 12
}
