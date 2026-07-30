package main

import "fmt"

type Buffer struct {
	buf     []byte
	initial [64]byte
	/*...*/
}

func (b *Buffer) Len() int {
	if b.buf == nil {
		return 0
	}

	return len(b.buf)
}

func (b *Buffer) Write(data []byte) {
	b.Grow(len(data))
	b.buf = append(b.buf, data...)
}

// Grow expands the buffer's capacity, if necessary,
// to guarantee space for another n bytes. [...]
func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		b.buf = b.initial[:0]
	}

	if len(b.buf)+n > cap(b.buf) {
		buf := make([]byte, b.Len(), 2*cap(b.buf)+n)
		copy(buf, b.buf)
		b.buf = buf
	}
}

func main() {
	var b Buffer

	fmt.Println("Length before writing:", b.Len()) // 0

	b.Write([]byte("hello"))
	fmt.Println("Length after 'hello':", b.Len()) // 5

	b.Write([]byte(" world!"))
	fmt.Println("Length after adding more:", b.Len()) // 12

	// Add a large chunk to force a reallocation via Grow().
	largeData := make([]byte, 100)
	for i := range largeData {
		largeData[i] = 'A'
	}

	b.Write(largeData)
	fmt.Println("Length after large chunk:", b.Len()) // 112
}
