package _func

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func MultipleReturnValues() {
	// 多个返回值
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
