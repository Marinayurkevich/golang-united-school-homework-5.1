package square

import "math"
	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (x, y Point) End() Point {
	x=x+a
	y=y-a
	return x, y
}

func (a Square) Area() uint {
	return math.Sqrt(a)
}

func (a Square) Perimeter() uint {
	return a*4
}
