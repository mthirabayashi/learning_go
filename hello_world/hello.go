package main

import "fmt"
import "github.com/mthirabayashi/initial_go_practice/export_example"

func main() {
  fmt.Printf("hello, world\n")
  // print_nums.print_nums() - cannot call because lowercase func are not exported
  nums.Testprintnums()
}
