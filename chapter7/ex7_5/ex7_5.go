// Exercise 7.5: The LimitReader function in the io package accepts an io.Reader r and
// a number of bytes n, and returns another Reader that reads from r but reports an end-of-line
// condition after n bytes. Implement it.
// func LimitReader(r io.Reader, n int64) io.Reader

package main

import "io"

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
	//
}
