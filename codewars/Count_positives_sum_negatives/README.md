# Count of positives / sum of negatives
Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.
Example

For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].


### [Link to Codewars challenge](https://www.codewars.com/kata/576bb71bbbcf0951d5000044)

## Solving
First, the function checks if the slice is empty and returns an empty result if so.  
If the slice has elements, it iterates through each value: if the number is negative, it adds it to the running sum; if positive, it increments the count of positive numbers. Zero values are ignored.

```go
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


```
## What could I have done better?
I believe that for a begginer, this code works really well.
I have also learned how to use *for-range* with Golang.  
I'm happy with the final result. I believe the code is readable and effective.