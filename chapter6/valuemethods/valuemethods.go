package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance        // method value
	fmt.Println(distanceFromP(q))      // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)

	scaleP := p.ScaleBy // method value
	scaleP(2)           // p becomes (2, 4)
	scaleP(3)           // then (6, 12)
	scaleP(10)          // then (60, 120)

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2, 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
}
