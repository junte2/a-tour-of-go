package main

import (
  "fmt"
  "math"
)

func main() {
  // Error sample
  // fmt.Println(math.pi)
  //
  // # command-line-arguments
  // ./exported-nameds.go:9:15: cannot refer to unexported name math.pi
  // ./exported-nameds.go:9:15: undefined: math.pi
  fmt.Println(math.Pi)
}
