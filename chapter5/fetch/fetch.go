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

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, n, err
}
