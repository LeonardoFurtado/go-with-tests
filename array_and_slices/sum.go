package array_and_slices

func Sum(arr []int) int {
	var result int
	for _, number := range arr {
		result += number
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) != 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}

	}

	return sums
}
