package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main() {
	path := Path{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	fmt.Println("Initial path:", path)

	// Addition: (1,1)+(5,5) → (6,6) и т.д.
	path.TranslateBy(Point{5, 5}, true)
	fmt.Println("After addition:", path)

	// Subtraction: (6,6)-(2,2) → (4,4) и т.д.
	path.TranslateBy(Point{2, 2}, false)
	fmt.Println("After subtraction:", path)
}
