package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(i int) int {
	if i < 1 {
		return 1
	}

	return i * factorial(i-1)
}
