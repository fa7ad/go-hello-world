package main

import "fmt"

func fib(n int) int {
  if n < 2 {
    return 1
  }
  return fib(n-1) + fib(n-2)
}

func main() {
  // 1 1 2 3 5 8 13 21 34 55 |89|
  // 0 1 2 3 4 5 6  7  8  9  |10|
  fmt.Println("10th Fibonacci number:", fib(10))
}