package main

import "fmt"

func main() {
  shorthand()
  non_shorthand()
}

func shorthand() {
  a := 10
  b := "Test String"
  c := 3.14
  d := true

  fmt.Printf("%v \n", a)
  fmt.Printf("%v \n", b)
  fmt.Printf("%v \n", c)
  fmt.Printf("%v \n\n", d)
}

func non_shorthand() {
  var a int
  var b string
  var c float64
  var d bool

  a = 10
  b = "Test String"
  c = 3.14
  d = true

  fmt.Printf("%v \n", a)
  fmt.Printf("%v \n", b)
  fmt.Printf("%v \n", c)
  fmt.Printf("%v \n", d)
}
