// Exercise 7.4: The strings.NewReader function returns a value that
// satisfies the io.Reader interface (and others) by reading from its argument, a string.
// Implement a simple version of NewReader yourself,
// and use it to make the HTML parser ($5.2) take input from a string.

package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	s string
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &StringReader{s}
}

func main() {
	reader := NewReader("Hello, World!")
	buf := make([]byte, 5)

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}

		fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
	}
}
