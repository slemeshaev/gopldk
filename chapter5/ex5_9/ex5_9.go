// Exercise 5.9: Write a function expand(s string, f func(string) string) string that
// replaces each substring "$foo" within s by the text returned by f("foo").
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`\$\w+`)

func main() {
	s, err := bufio.NewReader(os.Stdin).ReadString('.')
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", os.Args[0], err)
	}

	s = expand(s, strings.ToUpper)
	fmt.Fprintln(os.Stdout, s)
}

func expand(s string, f func(string) string) string {
	wrapper := func(s string) string {
		s = s[1:]
		return f(s)
	}

	return regex.ReplaceAllStringFunc(s, wrapper)
}
