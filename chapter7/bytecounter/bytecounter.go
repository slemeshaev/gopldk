package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c) // "5", = len("Hello")
	c = 0          // reset the counter

	var name = "Stanislav"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "16", = len("Hello, Stanislav")
}
