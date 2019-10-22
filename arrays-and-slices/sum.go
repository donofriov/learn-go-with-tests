package sum

// Sum takes in an array of numbers and returns the su,
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
