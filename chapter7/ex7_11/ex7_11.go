// Exercise 7.11: Add additional handlers so that clients can create, read, update, and delete database entries.
// For example, a request of the form /update?item=socks&price=6 will update the price of an item in the inventory
// and report an error if the item does not exist of if the price is invalid.
// (Warning: this change introduces concurrent variable updates.)

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
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
	mux.HandleFunc("/price", db.read)
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
	for item, price := range db.items {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.items[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Parse price: %v.\n", err)
		return
	}

	if _, exist := db.items[item]; exist {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Item %s already exists.\n", item)
	} else {
		db.items[item] = dollars(price)
		fmt.Fprintf(w, "Successfully created item. %s: %s", item, db.items[item])
	}
}

func (db *database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.items[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Parse price: %v.\n", err)
		return
	}

	if _, ok := db.items[item]; ok {
		db.items[item] = dollars(price)
		fmt.Fprintf(w, "Successfully updated item. %s: %s\n", item, db.items[item])
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db.items[item]; ok {
		delete(db.items, item)
		fmt.Fprintf(w, "Successfully deleted item %s\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q", item)
	}
}
