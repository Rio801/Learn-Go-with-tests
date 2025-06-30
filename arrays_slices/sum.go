package arraysslices

func Sum(numbers []int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}

// varidac params
func SumAll(nums ...[]int) []int {
	var sums []int
	for _, num := range nums {
		sums = append(sums, Sum(num))
	}
	return sums
}
