// Exercise 5.11: The instructor of the linear algebra course decides that
// calculus is now a prerequisite. Extend the topoSort function to report cycles.
package main

import (
	"fmt"
	"os"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms":           {"data structures"},
	"calculus":             {"linear algebra"},
	"linear algebra":       {"calculus"},        // circle
	"intro to programming": {"data structures"}, // another circle

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	sortReqs, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	for i, course := range sortReqs {
		fmt.Fprintf(os.Stdout, "%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for key := range m {
		path, circle := detectCircle(key, nil, m)
		if circle {
			return nil, fmt.Errorf("Circle detect: %s", strings.Join(path, " => "))
		}
		visitAll([]string{key})
	}

	return order, nil
}

func detectCircle(key string, path []string, m map[string][]string) ([]string, bool) {
	// Implementation
	return nil, false
}
