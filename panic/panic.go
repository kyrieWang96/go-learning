package main

import "fmt"

func main() {
	defer fmt.Println("test A")

	defer fmt.Println("test B")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()

	panic("hahahahaha")
}
