package main

import (
	"fmt"
	"math"
)

// Create the interface with
// two methods - area and perim
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// implement the methods for rects
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return r.width*2 + r.height*2
}

// implement the methods for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// create a measure function that accepts
// the interface as its type. Because there
// are implementations of both perim and area
// for circle and rect types, either type can be
// passed to this function. Each "g" calls its
// respective geometry function
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func newRect(width, height float64) *rect {
	r := rect{width: width, height: height}
	return &r
}

func newCircle(radius float64) *circle {
	c := circle{radius: radius}
	return &c
}

func main() {
	r := newRect(3, 4)
	c := newCircle(5)

	measure(r)
	measure(c)
}
