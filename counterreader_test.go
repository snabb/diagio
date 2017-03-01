package diagio_test

import (
	"bytes"
	"fmt"
	"github.com/snabb/diagio"
	"io"
	"io/ioutil"
)

func ExampleCounterReader() {
	input1 := bytes.NewBufferString("hello world")
	input2 := diagio.NewCounterReader(input1)
	io.Copy(ioutil.Discard, input2)
	fmt.Println(input2.Count())
	// Output:
	// 11
}
