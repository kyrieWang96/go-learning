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

// Fib /* 斐波那契数列：递归 */
func Fib(n int) int {
	// 终止条件 f(1) = 0, f(2) = 1
	if n == 1 || n == 2 {
		return n - 1
	}
	// 递归调用 f(n) = f(n-1) + f(n-2)
	res := Fib(n-1) + Fib(n-2)
	// 返回结果 f(n)
	return res
}
