package main

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

type EquilateralTriangle struct {
	Base   float64
	Height float64
}

func (e EquilateralTriangle) Area() float64 {
	return (e.Height * e.Base) / 2
}

func (e EquilateralTriangle) Perimeter() float64 {
	return 3 * e.Base
}
