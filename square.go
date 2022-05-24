package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	// implement me
	var a Point
	a.x = s.start.x + int(s.a)
	a.y = s.start.y + int(s.a)
	return a
	// or return Point{s.start.x + int(s.a), s.start.y + int(s.a)}
}

func (s Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	// implement me
	return s.a * 4
}
