package shapes

type Shape interface {
	Area() int
}

type Rectangle struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (rect Rectangle) Area() int {
	area := (rect.length * rect.breadth)
	return area
}

func (sq square) Area() int {
	area := (sq.side * sq.side)
	return area
}
