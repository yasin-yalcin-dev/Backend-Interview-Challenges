/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package shapehierarchy

import "fmt"

// RunExample demonstrates the shape hierarchy implementation
func RunExample() {
	// Create different shapes
	circle, err := NewCircle(5.0)
	if err != nil {
		fmt.Printf("Error creating circle: %v\n", err)
		return
	}

	rectangle, err := NewRectangle(4.0, 6.0)
	if err != nil {
		fmt.Printf("Error creating rectangle: %v\n", err)
		return
	}

	triangle, err := NewTriangle(3.0, 4.0, 5.0)
	if err != nil {
		fmt.Printf("Error creating triangle: %v\n", err)
		return
	}

	// Store them in a slice of Shape interface to demonstrate polymorphism
	shapes := []Shape{circle, rectangle, triangle}

	// Print information about each shape
	fmt.Println("Shape Hierarchy Example:")
	fmt.Println("------------------------")

	// Iterate and calculate areas and perimeters polymorphically
	for _, shape := range shapes {
		fmt.Printf("%s - Area: %.2f, Perimeter: %.2f\n",
			shape.Name(), shape.Area(), shape.Perimeter())
	}
}
