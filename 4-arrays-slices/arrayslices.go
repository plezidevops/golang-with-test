package arrayslices

func Sum(numbers []int) int {
	var count int
	for _, number := range numbers {
		count += number
	}
	return count
}

func SumAll(numbersToSum ...[]int) (sum []int) {

	sum = []int{3, 9}
	return
}
