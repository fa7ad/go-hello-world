package main

import "fmt"

func half(num int) (halfNum int, even bool) {
  halfNum = num / 2
  even = num % 2 == 0
  return
}

func main() {
  one, evn := half(1)
  two, even := half(2)

  fmt.Println("n", "n/2", "even?")
  fmt.Println(1, one, evn)
  fmt.Println(2, two, even)
}
