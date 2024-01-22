package recursion

// 比如我们求阶乘 1 * 2 * 3 * 4 * 5 *...* N

// Fact 阶乘
func Fact(x int) int {
	if x == 0 {
		return 1
	}
	return x * Fact(x-1)
}

func Fact1(x int, a int) int {
	if x == 1 {
		return a
	}

	return Fact1(x-1, x*a)
}
