package main

import "fmt"

func aboutMe(name string, age int) (string, int) {
	return "I am " + name, age
}

func main() {
	bio, age := aboutMe("Shahid", 25)

	fmt.Println(bio, "and I am", age, "years old.")
}
