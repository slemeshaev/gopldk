package main

import "fmt"

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Add adds the nonnegative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func main() {
	var set IntSet

	set.Add(64)

	fmt.Println(set.Has(64))  // true
	fmt.Println(set.Has(200)) // false

	var s1 IntSet
	s1.Add(1)
	s1.Add(3)
	fmt.Println("s1 = ", s1)

	var s2 IntSet
	s2.Add(2)
	s2.Add(4)
	fmt.Println("s2 = ", s2)

	s1.UnionWith(&s2)
	fmt.Println("After union s1 =", s1)
}
