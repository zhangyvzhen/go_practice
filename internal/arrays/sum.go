package arrays

func Sum(numbers []int) (n int) {
	for _, i := range numbers {
		n += i
	}
	return
}
