package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(z float64) {
	p.X *= z
	p.Y *= z
}

type Path []Point

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

type ColoredPoint struct {
	Point
	Color string
}

func (p ColoredPoint) String() string {
	return fmt.Sprintf("%s point at %g, %g", p.Color, p.X, p.Y)
}

func main() {
	x := Point{1, 2}
	y := Point{4, 6}

	fmt.Println(x.Distance(y), "<<<")
	fmt.Println(Point.Distance(x, y), "<<<")

	a := ColoredPoint{x, "yellow"}
	b := ColoredPoint{y, "green"}
	fmt.Println(a.Distance(b.Point)) // 5

	path := Path{x, y, Point{5, 6}}
	fmt.Println(path.Distance()) // 6

	c := ColoredPoint{Point{1, 2}, "yellow"}
	c.X = 2
	c.Y = 3
	c.ScaleBy(3)
	fmt.Println(c)
}
