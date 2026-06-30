// Exercise 5.18: Without changing its behavior, rewrite the fetch function
// to use defer to close the writable file.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	filename, n, err := fetch("https://lessgo.ru")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Downloaded: %s, size: %d bytes\n", filename, n)
}

// Fetch downloads the URL and returns the name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "" || local == "." {
		local = "index.html"
	}

	if !strings.Contains(local, ".") {
		local = local + ".html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
