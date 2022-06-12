package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}


func (t *Triangle) hasSide() bool {
	return t.Side >= 0
}

func (t *Triangle) CalcPerimeter() float64 {
	if !t.hasSide() {
		return 0
	}
	return 3 * t.Side
}

func (t *Triangle) CalcArea() float64 {
	if !t.hasSide() {
		return 0
	}
	return math.Sqrt(3) * math.Pow(t.Side, 2) / 4
}
