// Exercise 5.3: Write a function to print the contents of all text nodes
// in an HTML document tree. Do not descend into <script> or <style> elements, since
// their contents are not visible in a web browser.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr
var stdin io.Reader = os.Stdin

func main() {
	doc, err := html.Parse(stdin)
	if err != nil {
		fmt.Fprintf(stderr, "ex5_3: %v\n", err)
	}

	for _, text := range textfy(nil, doc) {
		fmt.Fprintln(stdout, text)
	}
}

func textfy(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.TextNode {
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			for _, line := range strings.Split(n.Data, "\n") {
				if len(line) != 0 {
					texts = append(texts, line)
				}
			}
		}
	}

	texts = textfy(texts, n.FirstChild)
	return textfy(texts, n.NextSibling)
}
