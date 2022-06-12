package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r *Rectangle) isRect() bool {
	return r.Height > 0 && r.Weight > 0
}

func (r *Rectangle) CalcPerimeter() float64 {
	if !r.isRect() {
		return 0
	}
	return 2 * (r.Height + r.Height)
}

func (r *Rectangle) CalcArea() float64 {
	if !r.isRect() {
		return 0
	}
	return r.Height * r.Weight
}
