package flameutil

import "math"

type Point struct{ X, Y float64 }

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p Point) Copy() *Point {
	return &p
}

func (p Point) R2() float64 {
	return p.X*p.X + p.Y*p.Y
}

func (p Point) R() float64 {
	return math.Sqrt(p.R2())
}

func (p Point) Angle() float64 {
	return math.Atan2(p.X, p.Y)
}
