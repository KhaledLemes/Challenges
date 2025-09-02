package main

func CountPositivesSumNegatives(numbers []int) []int {
  var res []int
  var sum int
  var positives int

  if len(numbers) == 0 {
	return res
  }
  
  for _, v := range numbers {
	if v < 0 {
		sum += v
	} else if v > 0 {
		positives++
	}
  }

  res = append(res, positives, sum)
  return res
}
