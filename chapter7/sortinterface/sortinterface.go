package main

import "fmt"

type Interface interface {
	Len() int
	Less(i, j int) bool // i, j are indeces of sequence elements
	Swap(i, j int)
}

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// bubbleSort sorts any collection implementing Interface,
// without knowing the concrete undelying type in advance
func bubbleSort(data Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

func main() {
	fruits := StringSlice{"banana", "apple", "cherry", "date"}

	fmt.Println("Before sorting:", fruits)
	bubbleSort(fruits)
	fmt.Println("After sorting:", fruits)
}
