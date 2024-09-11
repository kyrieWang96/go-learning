package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//var data map[string]string
	//data = make(map[string]string, 0)
	//data["s"] = "123"
	//fmt.Println(data)
	//var sfMap = SafeMap{}

	go func() {
		for i := 0; i < 10000; i++ {
			SetGoSafeMap("key"+strconv.Itoa(i), i)
		}

	}()

	go func() {
		for j := 0; j < 10000; j++ {
			SetGoSafeMap("key"+strconv.Itoa(j), j+10)
		}

	}()

	time.Sleep(time.Second * 10)
	//wg.Wait()
	fmt.Println(GoSafeMap)
}
