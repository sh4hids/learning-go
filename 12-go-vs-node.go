package main

import (
  "fmt"
  "time"
)

func main() {
	sum := 0
	start := time.Now().UnixNano() / int64(time.Millisecond)
	for i := 0; i < 100000000; i++ {
		sum += i
	}
  end := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Println(end)
	fmt.Println(start)
	fmt.Println(end - start)
}
