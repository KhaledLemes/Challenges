# String Repeat
Write a function that accepts a non-negative integer n and a string s as parameters, and returns a string of s repeated exactly n times.

Examples (input -> output)

```
6, "I" -> "IIIIII"

5, "Hello" -> "HelloHelloHelloHelloHello" 
```

### [Link to Codewars challenge](https://www.codewars.com/kata/57a0e5c372292dd76d000d7e)

## Solving
I used a simple for loop that sums up the string "result" with the string to be repeated.  
I am aware that this method uses more memory, however in order to not use an existing library and give a simple solution for a simple problem, I thought this approach would be a good one.

```go
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

```
## What could I have done better?
I do not thing that are improvements for this code without the use of external libraries.  
After reading other people's results, most looked like mine, and I do not believe that copying the strings.Repeat() library function would be a fair way to solve the problem
