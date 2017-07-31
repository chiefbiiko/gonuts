package main

import "fmt"

// finds minimum value in a slice
func findMin(x []int) int {
  var y int
  prev := 0
  for i := 0; i < len(x); i++ {
    if x[i] < prev {
      y = prev
    }
    prev = x[i]
  }
  return y
}

func main() {
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97,9,17,
  }
  fmt.Println(findMin(x))
}