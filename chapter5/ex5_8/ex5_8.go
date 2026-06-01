// Exercise 5.8: Modify forEachNode so that the pre and post functions
// return a boolean result indicating whether to continue the traversal.
// Use it to write a function ElementByID with the following signature
// that finds the first HTML element with the specified id attribute.
// The function should stop the traversal as soon as a match is found.
// func ElementByID(doc *html.Node, id string) *html.Node

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s HTML_FILE ID\n", os.Args[0])
		return
	}

	filename := os.Args[1]
	id := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(id, file)

	// Continue...
}
