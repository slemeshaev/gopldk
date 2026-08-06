// Exercise 7.2: Write a function CountingWriter with the signature below that,
// given an io.Writer, returns a new Writer that wraps the original,
// and a pointer to an int64 variable that at any moment contains the number of bytes written to the new Writer.
// func CountingWriter(w io.Writer) (io.Writer, *int64)

package main

import (
	"fmt"
	"io"
	"os"
)

type CounterWriter struct {
	writer io.Writer
	count  int64
}

func (w *CounterWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CounterWriter{
		writer: w,
		count:  0,
	}

	return cw, &cw.count
}

func main() {
	writer, count := CountingWriter(os.Stdout)

	fmt.Fprintln(writer, "Hello, World!")
	fmt.Printf("Bytes written: %d\n", *count)
}
