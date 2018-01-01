package main

import "fmt"

func main() {
  for i := 1; i < 101; i++ {
    output := ""
    if i % 3 == 0 {
      output += "Fizz"
    }
    if i % 5 == 0 {
      output += "Buzz"
    }
    if i % 3 != 0 && i % 5 != 0 {
      fmt.Print(i)
    }
    fmt.Println(output)
  }
}