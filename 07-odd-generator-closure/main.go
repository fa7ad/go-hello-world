package main

import "fmt"

func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (n uint) {
    n = i
    i += 2
    return
  }
}

func main() {
  nextOdd := makeOddGenerator()
  n1 := nextOdd()
  n2 := nextOdd()
  n3 := nextOdd()

  fmt.Println(n1, n2, n3)
}