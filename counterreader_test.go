package diagio_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/snabb/diagio"
)

func ExampleCounterReader() {
	input1 := bytes.NewBufferString("hello world")
	input2 := diagio.NewCounterReader(input1)
	io.Copy(ioutil.Discard, input2)
	fmt.Println(input2.Count())
	// Output:
	// 11
}
