package main

import "fmt"

func main() {
  fmt.Print("Enter a temperature in Fahrenheit: ")
  
  var fahr float64
  fmt.Scanf("%f", &fahr)

  cel := (fahr - 32) * 5.0 / 9.0

  fmt.Printf("Temperature in Celsius: %.2f", cel)
}