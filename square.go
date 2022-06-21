package square

import "math"
	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (x1, y1 Point) End() Point {
	x1 = x+a
	y1 = y-a
	return x1, y1
}

func (a Square) Area() uint {
s =  math.Sqrt(a)
	return s
}

func (b Square) Perimeter() uint {
p = math.Pow(a,4)
	return p
}
