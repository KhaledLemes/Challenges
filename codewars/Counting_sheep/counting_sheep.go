package main
func CountSheeps(numbers []bool) int {
  var result int
  for _, present := range numbers {
    if present {
      result++
    }
  }
  return result
}