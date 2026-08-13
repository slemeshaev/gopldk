// Exercise 7.5: The LimitReader function in the io package accepts an io.Reader r and
// a number of bytes n, and returns another Reader that reads from r but reports an end-of-line
// condition after n bytes. Implement it.
// func LimitReader(r io.Reader, n int64) io.Reader

package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitedReader struct {
	reader io.Reader
	limit  int64
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.limit {
		p = p[0:r.limit]
	}

	n, err = r.reader.Read(p)
	r.limit -= int64(n)

	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	fullReader := strings.NewReader("Hello, World! This is a test string.")
	limitedReader := LimitReader(fullReader, 10)

	buf := make([]byte, 3)
	for {
		n, err := limitedReader.Read(buf)
		if err == io.EOF {
			fmt.Println("EOF reached")
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
	}
}
