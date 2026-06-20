// Exercise 5.16: Write a variadic version of strings.Join.
package main

import (
	"fmt"
	"strings"
)

func Join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString(strs[0])

	for i := 1; i < len(strs); i++ {
		builder.WriteString(sep)
		builder.WriteString(strs[i])
	}

	return builder.String()
}

func main() {
	result := Join("*", "Go", "is", "the", "best", "language")
	fmt.Println(result)
}
