package integer

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers{
		sum += number
	}
	return sum
}

