package slice

func Sum(numbers []int) int {
	var result int

	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(numbers[1:]))
		}
	}

	return result
}
