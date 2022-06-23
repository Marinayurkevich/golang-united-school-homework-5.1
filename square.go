package square

	
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (e Square) End() Point {
e:=Point(x,y)
p:=&e
p.x = x+Square.a
	return e
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (p Square) Perimeter() uint {
	return 4 * p.a
}
