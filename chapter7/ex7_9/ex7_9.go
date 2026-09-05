// Exercise 7.9: Use the html/template package (§4.6) to replace `printTracks` with a function
// that displays the tracks as an HTML table. Use the solution to the previous exercise to arrange
// that each click on a column head makes an HTTP request to sort the table.

package main

import (
	"fmt"
	"os"
	"sort"
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

func tracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
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

type byColumns struct {
	tracks  []*Track
	columns []less
}

func sortByColumns(t []*Track, f ...less) *byColumns {
	return &byColumns{
		tracks:  t,
		columns: f,
	}
}

func (x byColumns) Len() int {
	return len(x.tracks)
}

func (x byColumns) Swap(i, j int) {
	x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i]
}

func (x byColumns) Less(i, j int) bool {
	if len(x.columns) == 0 {
		return false
	}

	a, b := x.tracks[i], x.tracks[j]
	var k int

	for k = 0; k < len(x.columns)-1; k++ {
		f := x.columns[k]
		switch {
		case f(a, b):
			return true
		case f(b, a):
			return false
		}
	}

	return x.columns[k](a, b)
}

func useSortByColumns() []*Track {
	t := tracks()
	sort.Sort(sortByColumns(t, colTitle, colArtist))
	return t
}

func useSortStable() []*Track {
	t := tracks()
	sort.Stable(byArtist(t))
	sort.Stable(byTitle(t))
	return t
}

type ColumnKey int

const (
	ColTitle ColumnKey = iota
	ColArtist
	ColAlbum
	ColYear
	ColLength
)

var comparators = map[ColumnKey]less{
	ColTitle:  colTitle,
	ColArtist: colArtist,
	ColAlbum:  colAlbum,
	ColYear:   colYear,
	ColLength: colLength,
}

type ColumnHistory struct {
	order []ColumnKey
}

func (h *ColumnHistory) click(col ColumnKey) {
	newOrder := make([]ColumnKey, 0, len(h.order)+1)
	newOrder = append(newOrder, col)

	for _, c := range h.order {
		if c != col {
			newOrder = append(newOrder, c)
		}
	}

	h.order = newOrder
}

func (h *ColumnHistory) less() []less {
	fs := make([]less, len(h.order))
	for i, col := range h.order {
		fs[i] = comparators[col]
	}
	return fs
}

func sortByHistory(t []*Track, h *ColumnHistory) {
	sort.Sort(byColumns{tracks: t, columns: h.less()})
}

func main() {
	t := tracks()
	h := &ColumnHistory{}

	fmt.Println("Click on Artist")
	h.click(ColArtist)
	sortByHistory(t, h)
	printTracks(t)

	fmt.Println("\nClick on Title")
	h.click(ColTitle)
	sortByHistory(t, h)
	printTracks(t)

	fmt.Println("\nClick on Artist")
	h.click(ColArtist)
	sortByHistory(t, h)
	printTracks(t)
}
