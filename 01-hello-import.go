package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	input := os.Args[1]
	fmt.Println(input)
	num, _ := strconv.ParseFloat(input, 64)
	fmt.Println(num)
	fmt.Println("The square root of", input, "is", math.Sqrt(num))
}
