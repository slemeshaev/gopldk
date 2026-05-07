package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
		if len(os.Args) != 3 {
			usageDie()
		}

		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "N (%s) must be an in", os.Args[1])
			usageDie()
		}

		comic, err := getComic(n)
		if err != nil {
			log.Fatal("Error getting comic", err)
		}

		fmt.Println(comic)
	case "index":
		if len(os.Args) != 3 {
			usageDie()
		}

		err := index(os.Args[2])
		if err != nil {
			log.Fatal("Error serializing indexes", err)
		}
	case "search":
		// Implementation
	default:
		usageDie()
	}
}

func getComic(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %d: %s", n, resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}

	return comic, nil
}

func index(filename string) error {
	nworkers := 20
	done := make(chan int)

	comicChan, err := getComics(nworkers, done)
	if err != nil {
		return err
	}

	go func() {
		for range nworkers {
			<-done
		}

		close(done)
		close(comicChan)
	}()

	indexComicsDealer(comicChan, filename)

	return nil
}

func getComics(nworkers int, done chan int) (chan Comic, error) {
	max, err := getComicCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("max", max)

	comics := make(chan Comic, 5*nworkers)
	comicsNums := make(chan int, 1*nworkers)

	for range nworkers {
		go fetcher(comicsNums, comics, done)
	}

	go dispatcher(comicsNums, max)

	return comics, nil
}

func indexComicsDealer(comicChan chan Comic, filename string) {
	// Need to implement
}

func getComicCount() (int, error) {
	// Need to implement
	return 0, nil
}

func fetcher(comicNums chan int, comics chan Comic, done chan int) {
	// Need to implement
}

func dispatcher(comicNums chan int, max int) {
	// Need to implement
}
