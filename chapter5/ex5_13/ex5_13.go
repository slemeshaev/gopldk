// Exercise 5.13: Modify crawl to make local copies of the pages it finds,
// creating directories as necessary. Don't make copies of pages that
// come from a different domain. For example, if the original page comes
// from golang.org, save all files from there, but exclude ones from vimeo.com.
package main

import (
	"net/url"
	"os"
)

func breadthFirst(f func(item, domain string) []string, worklist []string) {
	seen := make(map[string]bool)
	for _, w := range worklist {
		url, err := url.Parse(w)
		if err != nil {
			continue
		}

		domain := url.Host

		subworklist := make([]string, 1)
		subworklist[0] = w

		for len(subworklist) > 0 {
			items := subworklist
			subworklist = nil

			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					subworklist = append(subworklist, f(item, domain)...)
				}
			}
		}
	}
}

func crawl(url, domain string) []string {
	// Implementation
	return []string{}
}

func savePage(rawurl, domain string) error {
	// Implementation
	return nil
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
