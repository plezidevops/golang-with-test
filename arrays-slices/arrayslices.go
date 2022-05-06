package arrayslices

func Sum(numbers [5]int) int {
	var count int
	for _, number := range numbers {
		count += number
	}
	return count
}
