/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package shapehierarchy

import (
	"errors"
	"math"
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

// Circle represents a circle shape
type Circle struct {
	radius float64
}

// NewCircle creates a new Circle with validation
func NewCircle(radius float64) (*Circle, error) {
	if radius <= 0 {
		return nil, errors.New("radius must be positive")
	}
	return &Circle{radius: radius}, nil
}

// Area calculates the area of the circle
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter calculates the perimeter (circumference) of the circle
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Name returns the shape name
func (c *Circle) Name() string {
	return "Circle"
}

// Rectangle represents a rectangle shape
type Rectangle struct {
	width  float64
	height float64
}

// NewRectangle creates a new Rectangle with validation
func NewRectangle(width, height float64) (*Rectangle, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("width and height must be positive")
	}
	return &Rectangle{
		width:  width,
		height: height,
	}, nil
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter calculates the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Name returns the shape name
func (r *Rectangle) Name() string {
	return "Rectangle"
}

// Triangle represents a triangle shape
type Triangle struct {
	sideA float64
	sideB float64
	sideC float64
}

// NewTriangle creates a new Triangle with validation
func NewTriangle(a, b, c float64) (*Triangle, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, errors.New("all sides must be positive")
	}

	// Triangle inequality theorem: sum of the lengths of any two sides must be greater than the length of the remaining side
	if a+b <= c || a+c <= b || b+c <= a {
		return nil, errors.New("invalid triangle: does not satisfy triangle inequality theorem")
	}

	return &Triangle{
		sideA: a,
		sideB: b,
		sideC: c,
	}, nil
}

// Area calculates the area of the triangle using Heron's formula
func (t *Triangle) Area() float64 {
	// Semi-perimeter
	s := t.Perimeter() / 2

	// Heron's formula
	return math.Sqrt(s * (s - t.sideA) * (s - t.sideB) * (s - t.sideC))
}

// Perimeter calculates the perimeter of the triangle
func (t *Triangle) Perimeter() float64 {
	return t.sideA + t.sideB + t.sideC
}

// Name returns the shape name
func (t *Triangle) Name() string {
	return "Triangle"
}
