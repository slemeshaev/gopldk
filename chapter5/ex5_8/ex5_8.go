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

	"golang.org/x/net/html"
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

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	n := ElementByID(doc, id)
	if n == nil {
		fmt.Fprintf(os.Stdout, "ID %s not found in %s\n", id, filename)
	} else {
		fmt.Fprintf(os.Stdout, "ID %s found in %s\n", id, filename)
		for _, a := range n.Attr {
			fmt.Fprintf(os.Stdout, "<%s> has '%s' element, value is '%s'\n",
				n.Data, a.Key, a.Val)
		}
	}
}

func ElementByID(n *html.Node, id string) *html.Node {
	if n == nil {
		return nil
	}

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}

		return true
	}

	return forEachElement(n, pre, nil)
}

func forEachElement(n *html.Node, pre, pos func(n *html.Node) bool) *html.Node {
	// Implementation
	return &html.Node{}
}
