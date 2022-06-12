package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

var (
	nilBoxErr     = errors.New("nil box")
	outOfBoudsErr = errors.New("index out of bounds")
	NoCircleFoundErr = errors.New("no circles in the box")
)

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	if shapesCapacity < 0 {
		return nil
	}

	return &box{
		shapes:         make([]Shape, 0, shapesCapacity),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {

	if nil == b {
		return fmt.Errorf("%w", nilBoxErr)
	}

	if b.shapesCapacity < len(b.shapes) {
		return fmt.Errorf("%w", outOfBoudsErr)
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

//
func (b *box) hasBoxAt(i int) bool {
	return i >= 0 && i < len(b.shapes)
}

//

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {

	if nil == b {
		return nil, fmt.Errorf("%w", nilBoxErr)
	}

	if !b.hasBoxAt(i) {
		return nil, fmt.Errorf("%w", outOfBoudsErr)
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (s Shape, err error) {
	s, err = b.GetByIndex(i)
	if nil != err {
		return
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, new Shape) (old Shape, err error) {
	old, err = b.GetByIndex(i)
	if nil != err {
		return
	}
	old, b.shapes[i] = b.shapes[i], new
	return
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (s float64) {
	for _, v := range b.shapes {
		s += v.CalcPerimeter()
	}
	return
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (s float64) {
	for _, v := range b.shapes {
		s += v.CalcArea()
	}
	return
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() (err error) {

	if nil == b {
		err = fmt.Errorf("%w", nilBoxErr)
		return
	}

	var flag bool
	for k, v := range b.shapes {
		switch v.(type) {
		case *Circle:
			flag = true
			if k < len(b.shapes) {
				_, err = b.ExtractByIndex(k)
			} else {
				b.shapes = b.shapes[:k-1]
			}
		}
	}
	if !flag {
		err = fmt.Errorf("%w", NoCircleFoundErr)
	}
	return
}
