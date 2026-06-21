// Exercise 5.17: Write a variadic function ElementsByTagName that,
// given an HTML node tree and zero or more names, returns all the elements that
// match one of those names. Here are two example calls:
// func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
package main

import (
	"golang.org/x/net/html"
)

func main() {
	// Implementation
}

// ElementsByTagName that given an HTML node tree and zero or moe names, returns all the elements that match one of those names
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if len(name) == 0 {
		return nil
	}

	if doc.Type == html.ElementNode {
		for _, tag := range name {
			if doc.Data == tag {
				nodes = append(nodes, doc)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementsByTagName(c, name...)...)
	}

	return nodes
}
