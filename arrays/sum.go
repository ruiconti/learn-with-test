package arrays

func Sum(numbers []int) (sum int) {
	// for i := 0; i < 5; i++ {
	// sum += numbers[i]
	for _, number := range numbers {
		sum += number
	}
	return
}
