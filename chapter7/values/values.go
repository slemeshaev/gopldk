package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var w io.Writer
	fmt.Println("1. type: ", reflect.TypeOf(w))
	fmt.Println("1. value: ", reflect.ValueOf(w))

	w = os.Stdout
	fmt.Println("2. type: ", reflect.TypeOf(w))
	fmt.Println("2. value: ", reflect.ValueOf(w))

	w = new(bytes.Buffer)
	fmt.Println("3. type: ", reflect.TypeOf(w))
	fmt.Println("3. value: ", reflect.ValueOf(w))

	w = nil
	fmt.Println("4. type: ", reflect.TypeOf(w))
	fmt.Println("4. value: ", reflect.ValueOf(w))
}
