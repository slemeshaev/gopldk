// Exercise 5.15: Write variadic functions max and min, analogous to sum.
// What should these functions do when called with no arguments?
// Write variants that require at least one argument.
package main

import "fmt"

func main() {
	fmt.Println(max())
	fmt.Println(max(3))
	fmt.Println(max(1, 2, 3, 4))

	fmt.Println(min())
	fmt.Println(min(3))
	fmt.Println(min(1, 2, 3, 4))
}

func min(vals ...int) int {
	if len(vals) == 0 {
		panic("No Args")
	}

	m := vals[0]

	for _, val := range vals {
		if m < val {
			m = val
		}
	}

	return m
}

func max(vals ...int) int {
	if len(vals) == 0 {
		panic("No Args")
	}

	m := vals[0]

	for _, val := range vals {
		if m > val {
			m = val
		}
	}

	return m
}
