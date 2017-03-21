package nums

import "fmt"

// print_nums will not be exported because it is lowercase
// Testprintnums will be exported because it is uppercase

func print_nums() {
  fmt.Printf("%d - %b - %x\n", 42, 42, 42)
  fmt.Printf("%#x\n", 42)
  fmt.Printf("%#X\n", 42)
  // %d - decimal
  // %b - binary
  // %x - hexadecimal
  // %#x - hexadecimal with 0x prepended
  // %#X - upper case hexadecimal with 0x prepended
}

func Testprintnums() {
  print_nums()
}
