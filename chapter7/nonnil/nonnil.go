package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	fmt.Println("=== Demonstrating the nil-interface trap ===")
	demoBug()

	fmt.Println("\n=== Fixed version ===")
	demoFixed()
}

// demoBug shows how out != nil can lie to us when a typed
// nil pointer gets wrapped inside an interface value.
func demoBug() {
	var buf *bytes.Buffer
	if debug {
		buf = nil // intentionally left nil to reproduce the bug
	}

	fmt.Printf("buf == nil: %v (type: %T)\n", buf == nil, buf)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Reason: out != nil was true, even though buf was nil!")
		}
	}()

	f(buf) // NOTE: subtly incorrect!
}

// f is the function from the textbook, unchanged
func f(out io.Writer) {
	fmt.Printf("out == nil inside f: %v\n", out == nil) // false! that's the catch
	if out != nil {
		out.Write([]byte("done!\n")) // panics here if out is really a nil pointer
	}
}

// demoFixed shows the correct way to handle this -
// only assign a real, initialized value to an interface variable
func demoFixed() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // correct: actually initialize it
	}

	fFixed(buf)
	if debug {
		fmt.Print("Buffer contents: ", buf.String())
	}
}

func fFixed(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
