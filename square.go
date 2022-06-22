package square

import "math"
	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (e Square) End() Point {
	return e.start+e.a
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (p Square) Perimeter() uint {
	return 4 * p.a
}
