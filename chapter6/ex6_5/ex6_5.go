// Exercise 6.5: The type of each word used by IntSet is uint64,
// but 64-bit arithmetic may be inefficient on a 32-bit platform.
// Modify the program to use the uint type, which is the most efficient
// unsigned integer type for the platform. Instead of dividing by 64,
// define a constant holding the effective size of uint in bits, 32 or 64.
// You can use the perhaps too-clever expression 32 << (^uint(0) >> 63) for this pupose.

package main

import (
	"bytes"
	"fmt"
)

const uintsize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Add adds the nonnegative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintsize, uint(x%uintsize)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintsize, uint(x%uintsize)
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

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}

	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DiffrenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}

	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		}
	}

	for len(s.words) > 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	if len(t.words) > len(s.words) {
		s.words = append(s.words, make([]uint, len(t.words)-len(s.words))...)
	}

	for i, tword := range t.words {
		s.words[i] ^= tword
	}

	for len(s.words) > 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}

func bitCount(x uint) int {
	if uintsize == 64 {
		x = x - ((x >> 1) & 0x5555555555555555)
		x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
		x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
		x = x + (x >> 8)
		x = x + (x >> 16)
		x = x + (x >> 32)
		return int(x & 0x7f)
	}

	// uintsize == 32
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	return int(x & 0x1f)
}

// Return the number of elements
func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		l += bitCount(word)
	}

	return l
}

// Remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/uintsize, uint(x%uintsize)
	if word < len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
}

// Remove all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Return a copy of the set
func (s *IntSet) Copy() *IntSet {
	c := new(IntSet)
	c.words = make([]uint, len(s.words))
	copy(c.words, s.words)
	return c
}

// AddAll addes all the non-negative values to the set
func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

// Elems returns a slice containing the elements of the set
func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		for j := 0; j < uintsize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, i*uintsize+j)
			}
		}
	}
	return
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	first := true

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < uintsize; j++ {
			if word&(1<<uint(j)) != 0 {
				if !first {
					buf.WriteByte(' ')
				}
				first = false
				fmt.Fprintf(&buf, "%d", uintsize*i+j)
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}

func main() {
	x := new(IntSet)
	x.AddAll(1, 144, 30)

	x.Elems()

	fmt.Println(x.String())
}
