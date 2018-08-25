package main

import "fmt"

const PI float64 = 3.1416

func addNum(x, y float64) float64 {
	return x + y
}

func main() {
	msg := "Total is"
	num1, num2 := 3.56, 6.780
	fmt.Println(msg, addNum(num1, num2))
	fmt.Println("The value of PI is", PI)
}
