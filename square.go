package square

import "math"
	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (e Point) End() Point {
	return e.Point+a
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (p Square) Perimeter() uint {
	return 4 * p.a
}
