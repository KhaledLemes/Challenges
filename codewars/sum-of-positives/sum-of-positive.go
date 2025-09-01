package main

import (
	p "fmt"
)

var numbers []int = []int{1,2,-6,-6,-8,2}

func main() {
	p.Println(PositiveSum(numbers))
	
}

func PositiveSum(numbers []int) int {
	sum := 0

	for i := 0; i < len(numbers); i++ {
		if (numbers[i] > 0) {
			sum += numbers[i]
		}
	}
	return sum
}
