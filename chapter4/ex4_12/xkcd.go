package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type WordIndex map[string]map[int]bool
type NumIndex map[int]Comic

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

func getComicCount() (int, error) {
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("can't get main page: %s", resp.Status)
	}

	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}

	return comic.Num, nil
}

func indexComicsDealer(comicChan chan Comic, filename string) {
	wordIndex, numIndex := indexComics(comicChan)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(wordIndex)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = encoder.Encode(numIndex)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func indexComics(comics chan Comic) (WordIndex, NumIndex) {
	wordIndex := make(WordIndex)
	numIndex := make(NumIndex)

	for comic := range comics {
		numIndex[comic.Num] = comic
		fmt.Printf("Dealing with comic %d\n", comic.Num)

		scanner := bufio.NewScanner(strings.NewReader(comic.Transcript))
		scanner.Split(ScanWords)

		for scanner.Scan() {
			token := strings.ToLower(scanner.Text())
			if _, ok := wordIndex[token]; !ok {
				wordIndex[token] = make(map[int]bool, 1)
			}
			wordIndex[token][comic.Num] = true
		}
	}

	return wordIndex, numIndex
}

func fetcher(comicNums chan int, comics chan Comic, done chan int) {
	// Need to implement
}

func dispatcher(comicNums chan int, max int) {
	// Need to implement
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Need to implement
	return 0, token, nil
}
