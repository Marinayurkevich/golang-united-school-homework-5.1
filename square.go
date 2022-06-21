package square

import "math"
	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (start Point) End() Point {
	End=start+a
	return End
}

func (s Square) Area() uint {
s =  math.Sqrt(a)
	return s
}

func (p Square) Perimeter() uint {
p = a*4
	return p
}
