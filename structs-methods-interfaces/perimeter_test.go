package geometry

import "testing"

// Suppose that we need some geometry code to calculate the perimeter of a rectangle given a height and width.
// We can write a Perimeter(width float64, height float64) function, where float64 is for floating-point numbers like 123.45.
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Now let's create a function called Area(width, height float64) which returns the area of a rectangle.
func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
