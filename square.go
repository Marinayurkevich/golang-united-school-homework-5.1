package square

	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (v Square) End() Point {
var e Point
var a int = int(v.a)
e.x = e.x + a
e.y = e.x - a
	return e
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (p Square) Perimeter() uint {
	return 4 * p.a
}
