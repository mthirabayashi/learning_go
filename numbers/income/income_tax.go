package main

import "os"
import "fmt"
import "strconv"

func main() {
  var amount int = strconv.Atoi(os.Args[1])
  // var amount int = strconv.ParseInt(os.Args[1],10,64)
  // fmt.Println(amount)
  // fmt.Println(os.Args[1])
  fmt.Println(amount)

}
