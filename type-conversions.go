package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4

  // Error sample
  // var f float64 = math.Sqrt((x*x + y*y))
  //
  // # command-line-arguments
  // ./type-conversions.go:10:34: cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
  var f float64 = math.Sqrt(float64(x*x + y*y))

  // Error sample
  //var z uint = uint(f)
  //
  // # command-line-arguments
  // ./type-conversions.go:16:7: cannot use f (type float64) as type uint in assignment
  var z uint = f

  fmt.Println(x, y, z)
}
