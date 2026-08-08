// Exercise 7.3: Write a String method for the *tree type in gopl.io/ch4/treesort (4.4)
// that reveals the sequence of values in the tree.

package main

import (
	"bytes"
	"fmt"
)

type Tree struct {
	value       int
	left, right *Tree
}

func (t *Tree) String() string {
	order := make([]int, 0)
	order = appendValues(order, t)
	if len(order) == 0 {
		return "[]"
	}

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v", order)

	return b.String()
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func main() {
	//
}
