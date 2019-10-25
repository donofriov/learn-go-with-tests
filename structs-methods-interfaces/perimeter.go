package geometry

// Perimeter takes a height and width and returns (h+w)*2
func Perimeter(height, width float64) float64 {
	return (height + width) * 2
}

// Area takes a height and width and returns h*w
func Area(height, width float64) float64 {
	return height * width
}
