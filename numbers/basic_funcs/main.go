package main

import "fmt"

func main() {
  a := 5
  b := 8
  fmt.Println(add(a,b))
  fmt.Println(subtract(a,b))
  // can successfully reassign values
  a = 2
  b = 16
  fmt.Println(multiply(a,b))
  c := 20.0
  d := 8.0
  fmt.Println(divide(c,d))
}

func add(x int, y int) int {
  return x + y
}
func subtract(x int, y int) int {
  return x - y
}
func multiply(x int, y int) int {
  return x * y
}
func divide(x float64, y float64) float64 {
  return x / y
}
