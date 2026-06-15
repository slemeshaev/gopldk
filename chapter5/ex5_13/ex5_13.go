// Exercise 5.13: Modify crawl to make local copies of the pages it finds,
// creating directories as necessary. Don't make copies of pages that
// come from a different domain. For example, if the original page comes
// from golang.org, save all files from there, but exclude ones from vimeo.com.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	fmt.Println(url)

	err := savePage(url, domain)
	if err != nil {
		log.Printf("Can't save URL \"%s\": %s", url, err)
	}

	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func savePage(rawurl, domain string) error {
	url, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("bad url: %s", err)
	}

	if domain != url.Host {
		return nil
	}

	dir := url.Host
	var filename string
	if filepath.Ext(url.Path) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = url.Path
	}

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}

	return err
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
