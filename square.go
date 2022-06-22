package square

	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (e Square) End() Point {
e.start:=Point{
x=x+e.a
y=y+e.a
}
	return e.start
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (p Square) Perimeter() uint {
	return 4 * p.a
}
