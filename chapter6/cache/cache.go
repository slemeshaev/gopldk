package main

import (
	"fmt"
	"sync"
)

// var (
// 	mu sync.Mutex // guards mapping
// 	mapping = make(map[string]string)
// )

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	// Add data
	cache.mapping["user1"] = "Alice"
	cache.mapping["user2"] = "Bob"
	cache.mapping["user3"] = "Charlie"

	// Search
	fmt.Println(Lookup("user1")) // Alice
	fmt.Println(Lookup("user2")) // Bob
	fmt.Println(Lookup("user3")) // Charlie
	fmt.Println(Lookup("user4")) // (empty line)
}
