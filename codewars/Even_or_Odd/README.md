# Even Or Odd
Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.  
It should return "Even" for 0

### [Link to Codewars challenge](https://www.codewars.com/kata/53da3dbb4a5168369a0000fe)

## Solving
This is already a well documented function.  
I have simply created an *if* statement that returns "Even" if number % 2 == 0 and "Odd" otherwise

```go
func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "even"
	}
	return "odd"
}
```
In python the same expression can be written in one line of code (or 2 if we are including the function declaration on the count):

```python
def even_or_odd(num):
    return "Even" if num %2 == 0 else "Odd"
```
## What could I have done better?
This approach is the most used and simplest one, in both languages, but I found better solutions than this one.  
One example was using the & operator (*binary and*).
```go
func EvenOrOdd(number int) string {
  return []string{"Even", "Odd"}[number & 1]
}
```
In this approach, the & operator will compare the number's value with the binary digits of one.<br>
If both compared digits are 1, it will return 1.<br>
Since 1 only has one binary digit which is a 1 on the end of it and every even binary number ends with 0, this function will always return "Even" (number[0]) if an even number is put on it. 

  00001101  (13) ////// 00000010 (2)<br>
 00000001  (1)  //////// 00000001 (1)<br>
_--------------------_ /////// _--------------------_  
  00000001  (1) /////// 00000000 (0)