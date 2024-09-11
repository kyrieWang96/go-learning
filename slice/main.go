package main

import "fmt"

func main() {
	//sliceA := make([]string, 0)
	//sliceA = append(sliceA, "A", "B", "C")
	//
	//sliceB := make([]string, 0)
	//sliceB = sliceA
	//sliceB[0] = "Z"
	//
	//fmt.Println(sliceA)
	//fmt.Println(sliceB)
	//
	//a := [2]string{"1", "2"}
	//b := [3]string{"1", "2"}
	//if reflect.DeepEqual(a, b) {
	//	fmt.Println("is ok ")
	//} else {
	//	fmt.Println("no ok")
	//}

	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{7, 8, 9}
	//// copy(slice1, slice2)
	//copy(slice2, slice1)
	//fmt.Println(slice2)

	//var sliceA []string
	//sliceB := make([]string, 0)
	//if sliceA == nil {
	//	fmt.Println("is nil")
	//}
	//if reflect.DeepEqual(sliceA, sliceB) {
	//	fmt.Println("is equal")
	//} else {
	//	fmt.Println("no equal")
	//}
	// A := make([]string, 0)
	//A = append(A, "A", "B", "C")
	//A := [3]string{"A", "B", "C"}
	//B := ChangeSlice(A)
	//fmt.Println(A)
	//fmt.Println(B)
	s := []string{"A", "B"}

	fmt.Println(AppendSlice(s))
	fmt.Println(s)
}

func ChangeSlice(A [3]string) [3]string {
	A[0] = "z"
	return A
}

func AppendSlice(slice []string) *[]string {
	slice = append(slice, "1", "2", "3")
	return &slice
}
