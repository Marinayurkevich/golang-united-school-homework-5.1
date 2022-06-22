package square

import "math"
	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver) End() Point {
	End = Square.start+a
	return End
}

func (receiver) Area() uint {
Area =  math.Sqrt(a)
	return Area
}

func (receiver) Perimeter() uint {
Perimeter = 4*Square.a
	return Perimeter
}
