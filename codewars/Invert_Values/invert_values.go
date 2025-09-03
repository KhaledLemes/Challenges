package main

func Invert(arr []int) []int {
	result := make([]int, len(arr)) //Result's lenght is always the same as the argument to avoid panic
	
	for i, v := range arr {
		result[i] = -v
	}
	return result
}