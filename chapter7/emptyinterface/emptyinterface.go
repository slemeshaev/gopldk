package main

import (
	"bytes"
	"fmt"
)

func main() {
	var any interface{}
	any = true
	fmt.Println(any)

	any = 12.34
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = map[string]int{"one": 1}
	fmt.Println(any)

	any = new(bytes.Buffer)
	fmt.Println(any)
}
