package main

func Summation(n int) int {
	var temp int
	if n <= 1 {
		return n
	}
	for i := 0; i <= n; i++ {
		temp += i
	}
    return temp
}