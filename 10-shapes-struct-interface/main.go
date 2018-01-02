package main

import (
  "fmt"
  "math"
)

func distance(x1, y1, x2, y2 float64) float64 {
  return math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))
}

type Shape interface {
  area() float64
  perimeter() float64
}

type Circle struct {
  x, y, r float64
}
func (c Circle) area() float64 {
  return math.Pi * c.r*c.r
}
func (c Circle) perimeter() float64 {
  return 2 * math.Pi * c.r
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}
func (r Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}
func (r Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return 2 * (l + w)
}

func totalDim (shapes ...Shape) (area, perimeter float64) {
  area, perimeter = 0.0, 0.0
  for _, v := range shapes {
    area += v.area()
    perimeter += v.perimeter()
  }
  return
}

func main() {
  circ := Circle{0, 0, 10}
  fmt.Println("Area of circle:", circ.area())
  fmt.Println("Perimeter of circle:", circ.perimeter())

  rect := Rectangle{0, 0, 5, 10}
  fmt.Println("Area of rectangle:", rect.area())
  fmt.Println("Area of rectangle:", rect.perimeter())

  ar, per := totalDim(circ, rect)
  fmt.Println("Combined area:", ar)
  fmt.Println("Combined perimeter:", per)
}