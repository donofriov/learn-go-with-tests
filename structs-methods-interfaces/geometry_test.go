package geometry

import "testing"

// Suppose that we need some geometry code to calculate the perimeter of a rectangle given a height and width.
// We can write a Perimeter(width float64, height float64) function, where float64 is for floating-point numbers like 123.45.
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Now let's create a function called Area(width, height float64) which returns the area of a rectangle.
// Our next requirement is to write an Area function for circles.
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
