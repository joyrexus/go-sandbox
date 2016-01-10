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

func (p Point) Add(q Point) Point {
	p.X += q.X
	p.Y += q.Y
	return p
}

func (p Point) Sub(q Point) Point {
	p.X -= q.X
	p.Y -= q.Y
	return p
}

func (p *Point) ScaleBy(q float64) {
	p.X *= q
	p.Y *= q
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) Path {
	var op func (p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
	return path
}

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func main() {
	x := Point{1, 2}
	y := Point{4, 6}

	fmt.Println(x.Add(y), "<<<") // 5, 8
	fmt.Println(x.Sub(y), "<<<") // 1, 2

	path := Path{x, y, Point{5, 6}}
	fmt.Println(path.Distance()) // 6

	t := path.TranslateBy(Point{1, 1}, true)
	fmt.Println(t)

	/*
	p := Point{1, 2}
	p.X = 2
	p.Y = 3
	p.ScaleBy(3)
	fmt.Println(p) // {6 9}
	*/
}
