// Exercise 5.7: Develop startElement and endElement into a general HTML pretty-printer.
// Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use
// short forms like <img/> instead of <img></img> when an element has no chi ldren. Write a
// test to ensure that the output can be parsed successf ully. (See Chapter 11.)
package main

import (
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	prettify(doc)
}

func prettify(n *html.Node) {
	start(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		prettify(c)
	}
	end(n)
}

func start(n *html.Node) {
	// Implementation
}

func end(n *html.Node) {
	// Implementation
}
