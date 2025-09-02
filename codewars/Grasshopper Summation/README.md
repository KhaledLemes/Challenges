# Grasshopper Summation
Write a program that finds the summation of every number from 1 to num (both inclusive). The number will always be a positive integer greater than 0. Your function only needs to return the result, what is shown between parentheses in the example below is how you reach that result and it's not part of it, see the sample tests.

For example (Input -> Output):
```
2 -> 3 (1 + 2)
8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)
```

### [Link to Codewars challenge](https://www.codewars.com/kata/55d24f55d7dd296eb9000030)

## Solving
The function initializes a temporary variable to accumulate the result. If n is equal to or lesser than 1, it returns n.<br>Otherwise, it iterates from 0 to n, adding each value to the accumulator, and finally returns the sum.

```go
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

```
## What could I have done better?
Although my loop-based solution works correctly, I later realized (after checking other people’s solutions) that there is a direct mathematical formula for summing numbers from 1 to n: n * (n + 1) / 2.
This approach is much more efficient, as it eliminates the need for iteration.