package golang_united_school_homework

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) hasRadius() bool {
	return c.Radius > 0
}

func (c *Circle) CalcPerimeter() float64 {
	if !c.hasRadius() {
		return 0
	}
	return 2 * math.Pi * c.Radius
}

func (c *Circle) CalcArea() float64 {
	if !c.hasRadius() {
		return 0
	}
	return math.Pi * math.Pow(c.Radius, 2)
}
