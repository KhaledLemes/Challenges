package main

func CountBits(n uint) int {
	result := 0
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			result++
		}
	}
	return result
}