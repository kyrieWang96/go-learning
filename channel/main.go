package main

import (
	"fmt"
	"strconv"
	"sync"
)

var ch1, ch2 = make(chan int), make(chan int)
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go PrintA()
	go PrintB()
	ch1 <- 1
	wg.Wait()

	fmt.Println("DONE")
}

func PrintA() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-ch1
		fmt.Println("A" + strconv.Itoa(i))
		ch2 <- 1
	}
	<-ch1
}

func PrintB() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		<-ch2
		fmt.Println("B")
		ch1 <- 1
	}
}
