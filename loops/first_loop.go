package main
import "fmt"
func main() {
  fmt.Printf("dec \t binary \t hex\n")
  for i := 0; i < 100; i++ {
    fmt.Printf("%d \t %b \t %#X \t %q\n", i, i, i, i)
  }
}
