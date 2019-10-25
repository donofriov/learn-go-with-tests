package geometry

// Rectangle struct has a height and a width
type Rectangle struct {
	Height float64
	Width  float64
}

// Perimeter takes a height and width and returns (h+w)*2
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

// Area takes a height and width and returns h*w
func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
