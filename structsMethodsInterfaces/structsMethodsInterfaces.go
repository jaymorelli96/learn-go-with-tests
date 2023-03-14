package structsMethodsInterfaces

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radious float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radious, 2)
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}
