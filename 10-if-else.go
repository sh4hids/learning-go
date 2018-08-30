package main

import "fmt"

func main() {
	num := 60

	if num % 2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

  if num % 4 == 0 {
    fmt.Println(num, "is divisible by 4.")
  }

  if num < 10 {
    fmt.Println(num, "is less than 10.")
  } else if num > 10 && num < 50 {
    fmt.Println(num, "is in between 10 and 50.")
  } else {
    fmt.Println(num, "is greater than 50.")
  }

  if name := "Shahid"; name == "Shahid" {
    fmt.Println("The code is written by", name)
  } else {
    fmt.Println("The code is written by someone else.")
  }

}
