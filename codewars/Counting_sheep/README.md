# Counting Sheep
Consider an array/list of sheep where some sheep may be missing from their place. We need a function that counts the number of sheep present in the array (true means present).

For example,
```py
[True,  True,  True,  False,
  True,  True,  True,  True ,
  True,  False, True,  False,
  True,  False, False, True ,
  True,  True,  True,  True ,
  False, False, True,  True]
```
The correct answer would be 17.

Don't forget to check for bad values like null/undefined

Please note: not needed in languages that treats them for you like Golang  
This is the reason I did an answer in Java as well

### [Link to Codewars challenge](https://www.codewars.com/kata/54edbc7200b811e956000556)

## Solving
I simply wrote a for loop that increases the counter for each item that is true.

```go
func CountSheeps(numbers []bool) int {
  var result int
  for _, present := range numbers {
    if present {
      result++
    }
  }
  return result
}
```
I also wrote an example on a language that does not treat *null* values, in this case, Java:  
```java
    public static int countSheeps(Boolean[] arrayOfSheeps) {
      int count = 0;
      for (Boolean present : arrayOfSheeps) {
        if (present != null && present) {
          count++;
        }
      }
      return count;
    }
```
## What could I have done better?