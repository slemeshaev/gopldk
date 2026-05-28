package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// forEachNode calls the functions pre(x) and post(x)
// for each node x in the tree rooted at n. Both functions are optional.
// pre is called before the childre are visited (preorder)
// and post is called after (postorder).
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

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	links := []string{}

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
	}

	forEachNode(doc, pre, nil)

	fmt.Printf("Found %d links on %s:\n", len(links), url)
	for i, link := range links {
		if i >= 20 {
			fmt.Printf("... and %d more\n", len(links)-20)
			break
		}

		fmt.Printf("  %d. %s\n", i+1, link)
	}
}
