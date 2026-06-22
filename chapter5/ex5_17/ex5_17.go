// Exercise 5.17: Write a variadic function ElementsByTagName that,
// given an HTML node tree and zero or more names, returns all the elements that
// match one of those names. Here are two example calls:
// func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	file, err := os.Open("page.html")
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(fmt.Sprintf("HTML parsing error: %v", err))
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	fmt.Println("Images:")
	for i, node := range images {
		for _, attr := range node.Attr {
			if attr.Key == "src" {
				fmt.Printf("  [%d] src=%s\n", i, attr.Val)
			}
		}
	}

	fmt.Println("\nHeadings:")
	for i, node := range headings {
		fmt.Printf("  [%d] tag=%s, text=%q\n", i, node.Data, getNodeText(node))
	}
}

func getNodeText(n *html.Node) string {
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
		}
	}
	return text
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
