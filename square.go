package main

import "fmt"

func main() {
	testSquare := Square{
		start: Point{5, 5},
		a:     5,
	}
	testSquare2 := NewSquare(10, 10, 10)
	fmt.Println(testSquare)
	fmt.Println(testSquare.End())
	fmt.Println(testSquare.Area())
	fmt.Println(testSquare.Perimeter())
	fmt.Println()
	fmt.Println(testSquare2)
	fmt.Println(testSquare2.End())
	fmt.Println(testSquare2.Area())
	fmt.Println(testSquare2.Perimeter())

}

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) String() string {
	return fmt.Sprintf("obgect type: %T\nstart point: (%v, %v)\nside length: %v\n", s, s.start.x, s.start.y, s.a)
}

func NewSquare(x, y int, a uint) Square {
	return Square{Point{x, y}, a}
}

func (s Square) End() Point {
	// implement me
	/*var a Point
	a.x = s.start.x + int(s.a)
	a.y = s.start.y + int(s.a)
	return a*/
	return Point{s.start.x + int(s.a), s.start.y + int(s.a)}
}

func (s Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	// implement me
	return s.a * 4
}
