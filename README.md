diagio
======

[![GoDoc](https://godoc.org/github.com/snabb/diagio?status.svg)](https://godoc.org/github.com/snabb/diagio)

The Go package diagio implements I/O wrappers which can be useful for
debugging and other diagnostics purposes.

CounterReader implements an io.Reader which keeps track of bytes read
from it.

CounterWriter implements an io.Writer which keeps track of bytes written
to it.

DiagReader implements an io.Reader which produces diagnostics output
(number of bytes, current position, error, hex dump of data) whenever
it is read from.

DiagWriter implements an io.Writer which produces diagnostics output
(number of bytes, current position, error, hex dump of data) whenever
it is written to.

Documentation:

https://godoc.org/github.com/snabb/diagio

Simple example:
```
	cw := diagio.NewCounterWriter(os.Stdout)
	fmt.Fprintln(cw, "hello world")
	fmt.Println(cw.Count())
	// Output:
	// hello world
	// 12
```

The Git repository is located at: https://github.com/snabb/diagio


License
-------

MIT
