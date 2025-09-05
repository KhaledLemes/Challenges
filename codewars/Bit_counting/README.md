# Bit Counting
Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

### [Link to Codewars challenge](https://www.codewars.com/kata/526571aae218b8ee490006f4)

## Solving
I created a loop that iterates through the 32 bits of a number.
It shifts the bits of a number to the right *i* times, and if the Least Significant Bit is 1, the counter increases.

```go
func CountBits(n uint) int {
	result := 0
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			result++
		}
	}
	return result
}
```
## What could I have done better?
After reading other people's solutions I came to the conclusion that my approach was not bad, but could have been better. Iterating 32 times may sacrifice performance way too much, because it would do it even in 2, 4, 8... bits long numbers.