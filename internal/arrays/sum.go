package arrays

func Sum(numbers []int) (n int) {
	for _, i := range numbers {
		n += i
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	//make 可以在创建切片的时候指定我们需要的长度和容量
	sums = make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return
}
