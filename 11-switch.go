package main

import (
	"fmt"
	"time"
)

func main() {
	num := 3

	fmt.Print(num, " in word is ")

	switch num {
	case 2:
		fmt.Println("two.")
	case 3:
		fmt.Println("three.")
	case 4:
		fmt.Println("four.")
	}

  today := time.Now().Weekday()

  switch today {
  case time.Friday, time.Saturday:
    fmt.Println("It's a weekend.")
  default:
    fmt.Println("It's a weekday.")
  }

  fmt.Println(time.Now().Weekday())

}
