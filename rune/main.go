package main

import "fmt"

func main() {
	str := "123ABC456"
	out := ReflectStr([]rune(str))
	fmt.Println(string(out))
}

func ReflectStr(in []rune) []rune {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}
