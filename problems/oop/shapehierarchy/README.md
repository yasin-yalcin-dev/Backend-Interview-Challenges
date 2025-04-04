# Shape Hierarchy

## Problem
Implement a polymorphic shape class hierarchy in Go. Create an abstract Shape interface and concrete implementations for Circle, Rectangle, and Triangle. Each shape should be able to calculate its area and perimeter.

## Requirements
1. Create a Shape interface with Area() and Perimeter() methods
2. Implement concrete types for Circle, Rectangle, and Triangle
3. Each shape should store its dimensions (radius for circle, sides for rectangle, etc.)
4. Implement proper validation for shape dimensions
5. Demonstrate polymorphism by using a slice of different shapes

## Examples
```go
// Create different shapes
circle := shapes.NewCircle(5.0)
rectangle := shapes.NewRectangle(4.0, 6.0)
triangle := shapes.NewTriangle(3.0, 4.0, 5.0)

// Store them in a slice of Shape interface
shapes := []shapes.Shape{circle, rectangle, triangle}

// Iterate and calculate areas and perimeters polymorphically
for _, shape := range shapes {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
}