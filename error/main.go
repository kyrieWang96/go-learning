package main

import (
	"fmt"
	"github.com/work/go-learning/error/pkg"
)

func main() {
	in := "1"
	list := []string{"2", "3"}
	_, err := existStrInList(in, list)
	if err != nil {
		if pkg.GetType(err) == pkg.ErrorNotExistData {
			fmt.Println("错误类型处理")
		}
	}
}

func existStrInList(in string, list []string) (bool, error) {
	var exit bool
	for _, val := range list {
		if in == val {
			exit = true
			return exit, nil
		}
	}
	return false, pkg.ErrorNotExistData.Newf("数组中不存在该字符串:", in)
}
