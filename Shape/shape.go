package main

import (
	"fmt"
	"math"
	"time"
)

type Shape interface {
	// Area 计算面积
	Area() float64
	// Perimeter 计算周长
	Perimeter() float64
}

type Rectangle struct {
	L int64
	W int64
}

type Circle struct {
	R int64
}

func (r Rectangle) Area() float64 {
	return float64(r.W * r.L)
}

func (r Rectangle) Perimeter() float64 {
	return float64(2*r.W + 2*r.L)
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.R) * float64(c.R)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * float64(c.R)
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f; Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{W: 10, L: 3}
	//printShapeInfo(rect)

	circle := Circle{
		R: 5,
	}

	inputList := make([]Shape, 0)
	inputList = append(inputList, rect, circle)

	//printShapeInfo(circle)

	shapes := make(chan Shape, 10)
	defer close(shapes)

	go cunSumShapeInfo(shapes)
	for i := 0; i < 5; i++ {
		go readShapeInfo(inputList, shapes)
	}

	time.Sleep(10 * time.Second)

}

func readShapeInfo(s []Shape, shapes chan<- Shape) {
	for _, val := range s {
		shapes <- val
	}
}

func cunSumShapeInfo(shapes <-chan Shape) {
	for shape := range shapes {
		printShapeInfo(shape)
	}
}
