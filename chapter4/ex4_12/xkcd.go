package main

import (
	"fmt"
	"os"
)

type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string
}

const usage = `xkcd get N
xkcd index OUTPUT_FILE
xkcd search INDEX_FILE QUERY`

func usageDie() {
	fmt.Println(usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "get":
		// Implementation
	case "index":
		// Implementation
	case "search":
		// Implementation
	default:
		usageDie()
	}
}
