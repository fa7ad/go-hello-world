package main

import "fmt"

func main() {
  fmt.Print("Enter a length in feet: ")
  
  var feet float64
  fmt.Scanf("%f", &feet)

  met := feet * 0.3048

  fmt.Println("Length in Meters:", met)
}