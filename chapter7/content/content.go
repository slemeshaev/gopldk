package main

import (
	"fmt"
	"io"
	"time"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resulution() (x, y int)
}

type Book struct {
	title    string
	authors  []string
	created  time.Time
	pages    int
	words    int
	pageSize int
}

func (b Book) Title() string {
	return b.title
}

func (b Book) Creators() []string {
	return b.authors
}

func (b Book) Created() time.Time {
	return b.created
}

func (b Book) Pages() int {
	return b.pages
}

func (b Book) Words() int {
	return b.words
}

func (b Book) PageSize() int {
	return b.pageSize
}

func PrintArtifact(a Artifact) {
	fmt.Printf("Title: %s\n", a.Title())
	fmt.Printf("Creators: %v\n", a.Creators())
	fmt.Printf("Created: %s\n", a.Created().Format("2006-01-02"))
}

func PrintText(t Text) {
	fmt.Printf("Pages: %d\n", t.Pages())
	fmt.Printf("Words: %d\n", t.Words())
	fmt.Printf("PageSize: %d\n", t.PageSize())
}

func main() {
	book := Book{
		title:    "The Go Programming Language",
		authors:  []string{"Donovan", "Kernighan"},
		created:  time.Date(2015, time.November, 1, 0, 0, 0, 0, time.UTC),
		pages:    380,
		words:    120000,
		pageSize: 400,
	}

	fmt.Println("=== Book ===")
	PrintArtifact(book)
	PrintText(book)
}
