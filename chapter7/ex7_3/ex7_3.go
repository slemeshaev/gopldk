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

// Sort sorts values in place.
func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(Tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func main() {
	// Example 1: Manually creating a tree
	root := &Tree{value: 5}
	root.left = &Tree{value: 3}
	root.right = &Tree{value: 7}
	root.left.left = &Tree{value: 1}
	root.left.right = &Tree{value: 4}
	root.right.right = &Tree{value: 9}

	fmt.Println("In-order traversal:", root.String()) // [1 3 4 5 7 9]

	// Example 2: Slice sorting
	values := []int{5, 3, 7, 1, 4, 9, 2, 8, 6}
	fmt.Println("Before Sort:", values)

	Sort(values)
	fmt.Println("After Sort:", values) // [1 2 3 4 5 6 7 8 9]

	// Example 3: Empty tree
	empty := &Tree{}
	fmt.Println("Empty tree:", empty.String()) // []
}
