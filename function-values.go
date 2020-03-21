package main

import (
  "fmt"
  "math"
)

func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

func main() {
  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println(hypot(5, 12))       // Sqrt(5*5 + 12*12) = 13

  fmt.Println(compute(hypot))     // Sqrt(3*3 + 4+4) = 5
  fmt.Println(compute(math.Pow))  // 3^4 = 81
}
