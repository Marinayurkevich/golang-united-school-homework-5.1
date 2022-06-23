package square

	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (v Square) End() Point {
var a int = int(v.a)
v.start += a
	return v.start
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (p Square) Perimeter() uint {
	return 4 * p.a
}
