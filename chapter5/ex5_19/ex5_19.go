// Exercise 5.19: Use panic and recover to write a function
// that contains no return statement yet returns a non-zero value.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "simple: %d\n", simple())
}

func simple() (s int) {
	defer func() {
		if r := recover(); r != nil {
			s = 1
		}
	}()
	panic("ouch!")
}
