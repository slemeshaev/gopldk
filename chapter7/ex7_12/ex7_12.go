// Exercise 7.12: Change the handler for /list to print its output as an HTML table, not text.
// You may find the html/template package ($4.6) useful.

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

func main() {
	db := newDatabase(map[string]dollars{"shoes": 50, "socks": 5})
	log.Fatal(http.ListenAndServe("localhost:8000", db.routes()))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// database guards the map with a mutex: every HTTP request is served
// in its own goroutine, so without synchronization we get a data race.
type database struct {
	mu    sync.RWMutex
	items map[string]dollars
}

func newDatabase(init map[string]dollars) *database {
	items := make(map[string]dollars, len(init))
	for item, price := range init {
		items[item] = price
	}

	return &database{items: items}
}

func (db *database) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/read", db.read)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)

	return mux
}

// get returns the price of item and whether the item exists
func (db *database) get(item string) (dollars, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	price, ok := db.items[item]
	return price, ok
}

// add stores a new item. It reports false if the item already exists.
// The check and the write happen under one lock, so two concurrent
// creates of the same item cannot both succeed.
func (db *database) add(item string, price dollars) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; ok {
		return false
	}

	db.items[item] = price
	return true
}

// set changes the price of an existing item. It reports false if the item does not exist.
func (db *database) set(item string, price dollars) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; !ok {
		return false
	}

	db.items[item] = price
	return true
}

// remove deletes item. It reports false if the item does not exist.
func (db *database) remove(item string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; !ok {
		return false
	}

	delete(db.items, item)
	return true
}

type entry struct {
	item  string
	price dollars
}

// snapshot returns all items sorted by name.
func (db *database) snapshot() []entry {
	db.mu.RLock()

	entries := make([]entry, 0, len(db.items))
	for item, price := range db.items {
		entries = append(entries, entry{item, price})
	}

	db.mu.RUnlock()
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].item < entries[j].item
	})

	return entries
}

// parseRequest extracts item from the request and, if needPrice is set, validates price too.
func parseRequest(req *http.Request, needPrice bool) (item string, price dollars, err error) {
	q := req.URL.Query()
	item = q.Get("item")
	if item == "" {
		return "", 0, errors.New("item is required")
	}

	if !needPrice {
		return item, 0, nil
	}

	f, err := strconv.ParseFloat(q.Get("price"), 32)
	if err != nil {
		return "", 0, fmt.Errorf("invalid price: %w", err)
	}

	if math.IsNaN(f) || math.IsInf(f, 0) || f < 0 {
		return "", 0, fmt.Errorf("invalid price %q: must be a finite non-negative number", q.Get("price"))
	}

	return item, dollars(f), nil
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for _, e := range db.snapshot() {
		fmt.Fprintf(w, "%s: %s\n", e.item, e.price)
	}
}

// price replies with the bare price of the item, like in the book's example.
func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item, _, err := parseRequest(req, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	price, ok := db.get(item)
	if !ok {
		http.Error(w, fmt.Sprintf("no such item: %q", item), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

// read replies with "item: price"
func (db *database) read(w http.ResponseWriter, req *http.Request) {
	item, _, err := parseRequest(req, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	price, ok := db.get(item)
	if !ok {
		http.Error(w, fmt.Sprintf("no such item: %q", item), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item, price, err := parseRequest(req, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !db.add(item, price) {
		http.Error(w, fmt.Sprintf("item %q already exists", item), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "created item %s: %s\n", item, price)
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item, price, err := parseRequest(req, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !db.set(item, price) {
		http.Error(w, fmt.Sprintf("no such item: %q", item), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "updated item %s: %s\n", item, price)
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item, _, err := parseRequest(req, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !db.remove(item) {
		http.Error(w, fmt.Sprintf("no such item: %q", item), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "deleted item %s\n", item)
}
