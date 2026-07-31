package main

import "fmt"

type Counter struct {
	n int
}

func NewCounter(start int) *Counter {
	return &Counter{n: start}
}

func (c *Counter) N() int {
	return c.n
}

func (c *Counter) Increment() {
	c.n++
}

func (c *Counter) Reset() {
	c.n = 0
}

func main() {
	number := NewCounter(5)

	fmt.Println("n = ", number.N())

	number.Increment()
	fmt.Println("Increment = ", number.N())

	number.Reset()
	fmt.Println("Reset = ", number.N())
}
