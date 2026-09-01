// Exercise 7.8: Many GUIs provide a table widget with a stateful multi-tier sort: the primary sort key is the most recently
// clicked column head, the secondary sort key is the second-most recently clicked column head, and so on.
// Define an implementation of sort.Interface for use by such a table. Compare that approach with repeated sorting using sort.Stable.

package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "-----", "-----")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush() // calculate column widths and print table
}

type byTitle []*Track

func (x byTitle) Len() int {
	return len(x)
}

func (x byTitle) Less(i, j int) bool {
	return x[i].Title < x[j].Title
}

func (x byTitle) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byYear []*Track

func (x byYear) Len() int {
	return len(x)
}

func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type less func(x, y *Track) bool

func colTitle(x, y *Track) bool {
	return x.Title < y.Title
}

func colArtist(x, y *Track) bool {
	return x.Artist < y.Artist
}

func colAlbum(x, y *Track) bool {
	return x.Album < y.Album
}

func colYear(x, y *Track) bool {
	return x.Year < y.Year
}

func colLength(x, y *Track) bool {
	return x.Length < y.Length
}

func main() {
	//
}
