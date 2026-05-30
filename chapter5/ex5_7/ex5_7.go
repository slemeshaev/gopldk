// Exercise 5.7: Develop startElement and endElement into a general HTML pretty-printer.
// Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use
// short forms like <img/> instead of <img></img> when an element has no chi ldren. Write a
// test to ensure that the output can be parsed successf ully. (See Chapter 11.)
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int = 0

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
	switch n.Type {
	case html.ElementNode:
		startElement(n)
	case html.TextNode:
		startText(n)
	case html.CommentNode:
		startComment(n)
	}
}

func startElement(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}

	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}

	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	name := n.Data

	fmt.Fprintf(os.Stdout, "%*s<%s%s%s\n", depth*2, "", name, attrStr, end)
	depth++
}

func startText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	fmt.Fprintf(os.Stdout, "%*s%s\n", depth*2, "", n.Data)
}

func startComment(n *html.Node) {
	fmt.Fprintf(os.Stdout, "<!--%s-->\n", n.Data)
}

func end(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		endElement(n)
	}
}

func endElement(n *html.Node) {
	depth--
	if n.FirstChild == nil {
		return
	}
	fmt.Fprintf(os.Stdout, "%*s</%s>\n", depth*2, "", n.Data)
}
