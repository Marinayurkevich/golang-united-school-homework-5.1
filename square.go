package square

import "math"
	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (t1 Point) End() Point {
	x1 := x+a
	y1 := y+a
	return t1.x1, t1.y1
}

func (s Square) Area() uint {
s =  math.Sqrt(a)
	return s
}

func (p Square) Perimeter() uint {
p = a*4
	return p
}
