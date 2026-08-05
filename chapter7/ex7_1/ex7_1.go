// Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanWords)

	var count int
	for s.Scan() {
		count++
	}

	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	var count int
	for s.Scan() {
		count++
	}

	*c += LineCounter(count)
	return count, nil
}

func main() {
	// WordCounter
	var b ByteCounter
	b.Write([]byte("Hello"))
	fmt.Println("ByteCounter:", b) // 5

	// WordCounter
	var w WordCounter
	w.Write([]byte("Hello world, how are you?"))
	fmt.Println("WordCounter:", w) // 5

	// LineCounter
	var l LineCounter
	l.Write([]byte("first line\nsecond line\nthird line"))
	fmt.Println("LineCounter:", l) // 3
}
