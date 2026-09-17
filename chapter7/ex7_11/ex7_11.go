// Exercise 7.11: Add additional handlers so that clients can create, read, update, and delete database entries.
// For example, a request of the form /update?item=socks&price=6 will update the price of an item in the inventory
// and report an error if the item does not exist of if the price is invalid.
// (Warning: this change introduces concurrent variable updates.)

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Parse price: %v.\n", err)
		return
	}

	if _, exist := db[item]; exist {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Item %s already exists.\n", item)
	} else {
		db[item] = dollars(price)
		fmt.Fprintf(w, "Successfully created item. %s: %s", item, db[item])
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
