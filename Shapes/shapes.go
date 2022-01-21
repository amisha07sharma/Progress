package shapes

type Rectangle struct {
	length  int
	breadth int
}

func (rect Rectangle) Area() int {
	area := (rect.length * rect.breadth)
	return area
}
