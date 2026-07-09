package main

import "fmt"

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	list := &IntList{1, &IntList{2, &IntList{3, nil}}}
	fmt.Println(list.Sum()) // 6
}
