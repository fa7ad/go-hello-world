package main

import "fmt"

func largest (nums ...int) (max int) {
  max = nums[0]
  for _, n := range nums {
    if n > max {
      max = n
    }
  }
  return
}

func main() {
  fmt.Println("Numbers:", 1, 10, 45, 2, 4, 6, 7, 9, 0)
  fmt.Println("Largest:", largest(1, 10, 45, 2, 4, 6, 7, 9, 0))
}