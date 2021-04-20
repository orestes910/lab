package main

import "fmt"

// rect will have methods
// implemented for its type
type rect struct {
	width, height int
}

//   receiver  name   return-type
func (r *rect) area() int {
	// access rect fields
	return r.width * r.height
}

// val version
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func newRect(w, h int) *rect {
	r := rect{width: w, height: h}
	return &r
}

func main() {
	// create a 10 by 5 rectangle
	r := newRect(10, 5)

	// a rect called r can call area
	// and perim, since they both have
	// "rect" as their receiver type
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// With val
	rv := *r
	fmt.Println("area: ", rv.area())
	fmt.Println("perim: ", rv.perim())
}
