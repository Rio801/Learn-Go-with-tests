package structsmethodsinterfaces

import "math"

type Rectangle struct {
	Height float64
	Weight float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base   float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Weight)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Weight
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (triangle Triangle) Area() float64 {
	return (triangle.Height * triangle.Base) / 2
}
