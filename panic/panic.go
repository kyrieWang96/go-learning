package main

import (
	"fmt"
	"time"
)

func getChannel(ch chan int) {

	fmt.Printf("接受数据: %d", <-ch)
	time.Sleep(time.Second * 5)
}

func main() {
	ch := make(chan int)

	ch <- 1
	go getChannel(ch)
	time.Sleep(time.Second * 5)

}

//func main() {
//	defer fmt.Println("test A")
//
//	defer fmt.Println("test B")
//
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(err.(string))
//		}
//	}()
//
//	panic("hahahahaha")
//}
