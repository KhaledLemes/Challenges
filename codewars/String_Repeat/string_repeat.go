package main

func RepeatStr(repetitions int, str string) string {
	var result string
	if repetitions < 1 {
		return ""
	}
	for i := 0; i < repetitions; i++ {
		result += str
	}
	return result
}