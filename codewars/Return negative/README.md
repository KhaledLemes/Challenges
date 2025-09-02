# Return Negative
In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?
Examples

```
MakeNegative(1)    // return -1
MakeNegative(-5)   // return -5
MakeNegative(0)    // return 0
```

Notes<br>
    - The number can be negative already, in which case no change is required.<br>
    - Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.

### [Link to Codewars challenge](https://www.codewars.com/kata/55685cd7ad70877c23000102)

## Solving
For this challenge, I implemented a simple if–else condition. If the number is already negative, the function just returns it. Otherwise, if the number is positive, it returns its negative version by multiplying it by –1.

```go
package main

func MakeNegative(num int) int {
	if num < 0 {
		return num
	} else {
		return 1*-num
	}
}