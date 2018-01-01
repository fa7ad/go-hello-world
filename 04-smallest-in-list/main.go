package main

import "fmt"

func main() {
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }
  sm := x[0]

  for _, v := range x {
    if v < sm {
      sm = v
    }
  }

  fmt.Println("Smallest number is:", sm)
}