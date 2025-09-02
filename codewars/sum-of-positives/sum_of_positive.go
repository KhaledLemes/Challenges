package main

//I opted to only let the function on this file
func PositiveSum(numbers []int) int {
	sum := 0

	for i := 0; i < len(numbers); i++ {
		if (numbers[i] > 0) {
			sum += numbers[i]
		}
	}
	return sum
}