package arrayslices

func Sum(numbers []int) int {
	var count int
	for _, number := range numbers {
		count += number
	}
	return count
}
