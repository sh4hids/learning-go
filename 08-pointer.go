package main

import "fmt"

func main()  {
  x := 15
  a := &x // point to memory address
  fmt.Println(*a); // read the value of memory address
}
