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
	"math"
	"testing"
)

// Helper function to compare float64 values with tolerance
func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestCircle(t *testing.T) {
	// Test valid circle
	circle, err := NewCircle(5.0)
	if err != nil {
		t.Errorf("Expected no error for valid circle, got: %v", err)
	}

	expectedArea := math.Pi * 25.0
	if !floatEquals(circle.Area(), expectedArea, 0.001) {
		t.Errorf("Circle area calculation incorrect. Got: %f, Expected: %f", circle.Area(), expectedArea)
	}

	expectedPerimeter := 2 * math.Pi * 5.0
	if !floatEquals(circle.Perimeter(), expectedPerimeter, 0.001) {
		t.Errorf("Circle perimeter calculation incorrect. Got: %f, Expected: %f", circle.Perimeter(), expectedPerimeter)
	}

	// Test invalid circle
	_, err = NewCircle(-1.0)
	if err == nil {
		t.Errorf("Expected error for negative radius, got none")
	}
}

func TestRectangle(t *testing.T) {
	// Test valid rectangle
	rectangle, err := NewRectangle(4.0, 5.0)
	if err != nil {
		t.Errorf("Expected no error for valid rectangle, got: %v", err)
	}

	expectedArea := 20.0
	if !floatEquals(rectangle.Area(), expectedArea, 0.001) {
		t.Errorf("Rectangle area calculation incorrect. Got: %f, Expected: %f", rectangle.Area(), expectedArea)
	}

	expectedPerimeter := 18.0
	if !floatEquals(rectangle.Perimeter(), expectedPerimeter, 0.001) {
		t.Errorf("Rectangle perimeter calculation incorrect. Got: %f, Expected: %f", rectangle.Perimeter(), expectedPerimeter)
	}

	// Test invalid rectangle
	_, err = NewRectangle(-4.0, 5.0)
	if err == nil {
		t.Errorf("Expected error for negative width, got none")
	}

	_, err = NewRectangle(4.0, -5.0)
	if err == nil {
		t.Errorf("Expected error for negative height, got none")
	}
}

func TestTriangle(t *testing.T) {
	// Test valid triangle
	triangle, err := NewTriangle(3.0, 4.0, 5.0)
	if err != nil {
		t.Errorf("Expected no error for valid triangle, got: %v", err)
	}

	expectedArea := 6.0
	if !floatEquals(triangle.Area(), expectedArea, 0.001) {
		t.Errorf("Triangle area calculation incorrect. Got: %f, Expected: %f", triangle.Area(), expectedArea)
	}

	expectedPerimeter := 12.0
	if !floatEquals(triangle.Perimeter(), expectedPerimeter, 0.001) {
		t.Errorf("Triangle perimeter calculation incorrect. Got: %f, Expected: %f", triangle.Perimeter(), expectedPerimeter)
	}

	// Test invalid triangle
	_, err = NewTriangle(-3.0, 4.0, 5.0)
	if err == nil {
		t.Errorf("Expected error for negative side, got none")
	}

	// Test triangle inequality violation
	_, err = NewTriangle(1.0, 2.0, 10.0)
	if err == nil {
		t.Errorf("Expected error for triangle inequality violation, got none")
	}
}

func TestPolymorphism(t *testing.T) {
	// Create shapes
	circle, _ := NewCircle(5.0)
	rectangle, _ := NewRectangle(4.0, 5.0)
	triangle, _ := NewTriangle(3.0, 4.0, 5.0)

	// Store in slice of Shape interface
	shapes := []Shape{circle, rectangle, triangle}

	expectedAreas := []float64{math.Pi * 25.0, 20.0, 6.0}
	expectedPerimeters := []float64{2 * math.Pi * 5.0, 18.0, 12.0}

	// Test polymorphic behavior
	for i, shape := range shapes {
		if !floatEquals(shape.Area(), expectedAreas[i], 0.001) {
			t.Errorf("%s area calculation incorrect. Got: %f, Expected: %f", shape.Name(), shape.Area(), expectedAreas[i])
		}

		if !floatEquals(shape.Perimeter(), expectedPerimeters[i], 0.001) {
			t.Errorf("%s perimeter calculation incorrect. Got: %f, Expected: %f", shape.Name(), shape.Perimeter(), expectedPerimeters[i])
		}
	}
}
