# Invert Values
Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

[1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
[1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
[] --> []

You can assume that all values are integers. Do not mutate the input array.


### [Link to Codewars challenge](https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad)

## Solving
To solve this challenge, I used a simple for...range loop to iterate through the array, adding each number to the result in its opposite form.

```go
func Invert(arr []int) []int {
	result := make([]int, len(arr)) //Result's lenght is always the same as the argument to avoid panic
	
	for i, v := range arr {
		result[i] = -v
	}
	return result
}
```
## What could I have done better?
After submitting the challenge, I checked other people's solutions.<br>
The most optimized approach was essentially the same as mine.