package main

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func PrintShapeInfo(s Shape) {
	println("Area:", s.Area())
	println("Perimeter:", s.Perimeter())
}

func TestInterfaceImplementation() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}

	PrintShapeInfo(r)
	PrintShapeInfo(c)
}
