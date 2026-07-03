package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// Example 1: HTML with one title
	htmlStr1 := `<html><head><title>Hello, World!</title></head><body></body></html>`
	doc, _ := html.Parse(strings.NewReader(htmlStr1))
	title, err := soleTitle(doc)
	fmt.Printf("Title: %q, Error: %v\n", title, err)

	// Example 2: HTML with two titles (triggers error via panic)
	htmlStr2 := `<html><head><title>First</title><title>Second</title></head><body></body></html>`
	doc2, _ := html.Parse(strings.NewReader(htmlStr2))
	title, err = soleTitle(doc2)
	fmt.Printf("Title: %q, Error: %v\n", title, err)

	// Example 3: HTML without title
	htmlStr3 := `<html><head></head><body></body></html>`
	doc3, _ := html.Parse(strings.NewReader(htmlStr3))
	title, err = soleTitle(doc3)
	fmt.Printf("Title: %q, Error: %v\n", title, err)
}

// soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()

	// Bail out recursion if we find more than one non-empty title
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
